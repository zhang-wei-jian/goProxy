## Go反向代理

只需要配置`config.json`文件要转发的地址即可。

```bash
go build -o goProxy.exe .\main.go
```

我启动这个项目是为了在本地的web应用设置了跨域，我需要转发到其他端口同时解决跨域。

### 使用方法：

1. 配置`config.json`文件：

```json
{
    "proxyIP": "要转发的地址",
    "localListenIP": "启动在本地的端口"
}
```

2. 启动`goProxy.exe`