
###  配置

配置文件名称:

> pxylogtail.yaml

配置文件格式:yaml

配置文件路径:

* .
* .\config
* $HOME/.pxylogtail

```
name: pxylogtail 
server:localhost                  # log server
serverPort:5514                   # log server port
logfiles:
  - name: test1.log               # log 文件名
    path: .\                      # log 文件路径   
    app:  test1                   # 应用名称
    host: myhost1                 # 应用主机名
    mail: fcx@pingxingyun.com     # mail
    version: 0.0.1                # 版本号
  - name: test2.log
    path: .\
    app:  test2
    host: myhost2
    mail: test@test.com
    version: 3.2.1
```

日志采集系统使用 https://github.com/ekanite/ekanite

日志协议： RFC5424 

只支持 TCP

#### ekanite 配置 

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