package main

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/crewjam/rfc5424"
	"github.com/hpcloud/tail"
	"github.com/spf13/viper"
)

const ConfigFile = "remotelogtail"

func startTail(logfile LogSetup, lfs *LogFiles) {
	var filename = logfile.getPath()
	fmt.Println("start tail:", filename)
	t, err := tail.TailFile(filename, tail.Config{Follow: true, Poll: true})
	if err != nil {
		fmt.Println("tail open file failed:", err)
	}
	for line := range t.Lines {
		fmt.Println(filename, "new line:", line.Text)
		m := rfc5424.Message{
			Priority:  rfc5424.Daemon | rfc5424.Info,
			Timestamp: time.Now(),
			Hostname:  logfile.Host,
			AppName:   logfile.App,
			Message:   []byte(line.Text + "\r\n"),
		}
		m.AddDatum(logfile.Mail, "Revision", logfile.Version)
		b, err := m.MarshalBinary()
		if err != nil {
			fmt.Println("log line error:", err)
		} else {
			lfs.WriteLog(b)
		}
	}
}

type LogFiles struct {
	Name       string
	Server     string
	ServerPort int
	Logfiles   []LogSetup
	conn       net.Conn
}

func (this LogFiles) getLogServer() (s string, err error) {
	if this.ServerPort != 0 && this.Server != "" {
		s = this.Server + ":" + strconv.Itoa(this.ServerPort)
	} else {
		err = errors.New("can't find server or server port")
	}
	return
}

func (this *LogFiles) connectLogServer() (err error) {
	fmt.Println("connect to log server")
	server, err := this.getLogServer()
	if err != nil || server == "" {
		fmt.Println("connect server err:", err)
		return
	}
	this.conn, err = net.Dial("tcp", server)
	if err != nil {
		fmt.Println("connect to log server error:", err)
		return
	}
	fmt.Println("connect to log server success")
	return
}

func (this *LogFiles) close() (err error) {
	err = this.conn.Close()
	return
}

func (this *LogFiles) WriteLog(b []byte) {
	lock.Lock()
	defer lock.Unlock()
	if this.conn != nil {
		this.conn.Write(b)
	} else {
		fmt.Println("connection not exist")
	}
}

type LogSetup struct {
	Name    string
	Path    string
	App     string
	Host    string
	Mail    string
	Version string
}

func (this LogSetup) getPath() string {
	return this.Path + this.Name
}
func (this LogSetup) String() string {
	return "Host:" + this.Host +
		";APP:" + this.App +
		"[@:" + this.Mail +
		"';v:" + this.Version +
		"];Path:" + this.Path + this.Name
}

var wg sync.WaitGroup
var lock *sync.RWMutex

func main() {
	fmt.Println("pxy log tail start!")
	lock = new(sync.RWMutex)
	defer fmt.Println("pxy log tail stop")
	viper.SetConfigName(ConfigFile) // name of config file (without extension)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("$HOME/.remotelogtail")
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("Config", viper.AllSettings())
	var lf LogFiles

	err = viper.Unmarshal(&lf)
	if err != nil {
		fmt.Println("unable to decode into struct, %v", err)
	}

	fmt.Printf("Log name:%s,Log server:%s:%d\r\n", lf.Name, lf.Server, lf.ServerPort)
	lf.connectLogServer()

	for key, val := range lf.Logfiles {
		fmt.Printf("log files key: %d; value: %s;\r\n", key, val)
		go startTail(val, &lf)
		wg.Add(1)
	}

	wg.Wait()
}
