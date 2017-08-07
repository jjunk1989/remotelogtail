#

Small soft for tail log to remote server ekanited.

Base on

![rfc5424](https://github.com/crewjam/rfc5424)
![tail](https://github.com/hpcloud/tail)
![viper](https://github.com/spf13/viper)

###  Config

Config name:

> remotelogtail.yaml

format:yaml

Config file find path:


* .
* .\config
* $HOME/.remotelogtail

```
name: remotelogtail 
server:localhost                  # log server
serverPort:5514                   # log server port
logfiles:
  - name: test1.log               # log file name
    path: .\                      # log path   
    app:  test1                   # App name
    host: myhost1                 # server host
    mail: test@test.com    		  # mail
    version: 0.0.1                # version
  - name: test2.log
    path: .\
    app:  test2
    host: myhost2
    mail: test@test.com
    version: 3.2.1
```

remote server https://github.com/ekanite/ekanite

#### ekanite config 

```
ekanited [options]
  -batchsize int
        Indexing batch size (default 300)
  -batchtime int
        Indexing batch timeout, in milliseconds (default 1000)
  -cpuprof string
        Where to write CPU profiling data. Not written if not set
  -datadir string
        Set data directory (default "/var/opt/ekanite")
  -diag string
        expvar and pprof bind address in the form host:port. If not set, not started (default "localhost:9951")
  -input string
        Message format of input (only syslog supported) (default "syslog")
  -maxpending int
        Maximum pending index events (default 1000)
  -memprof string
        Where to write memory profiling data. Not written if not set
  -numshards int
        Set number of shards per index (default 4)
  -query string
        TCP Bind address for query server in the form host:port. To disable setto empty string (default "localhost:9950")
  -queryhttp string
        TCP Bind address for http query server in the form host:port. To disableset to empty string (default "localhost:8080")
  -retention string
        Data retention period. Minimum is 24 hours (default "168h")
  -tcp string
        Syslog server TCP bind address in the form host:port. To disable set to empty string (default "localhost:5514")
  -tlskey string
        path to CA key file for TLS-enabled TCP server. If not set, TLS not activated
  -tlspem string
        path to CA PEM file for TLS-enabled TCP server. If not set, TLS not activated
  -udp string
        Syslog server UDP bind address in the form host:port. If not set, not started
```

```
ekanited -datadir .\logs
```