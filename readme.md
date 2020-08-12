## tcp 端口转发

### config.json

配置文件

- bindtcp 绑定本地地址
- totcp   转发地址

```
{
    "bindtcp" : "0.0.0.0:3388",
    "totcp" : "192.168.1.42:3389"
}
```


### 安装运行
```
md %GOPATH%\src\github.com\jiegemena
cd %GOPATH%\src\github.com\jiegemena
git clone ...
go build -o tcpProxy.1.0.exe
```