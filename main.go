package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhang-wei-jian/goProxy/hooks"
	"github.com/zhang-wei-jian/goProxy/tools"
)

var JSON tools.Config

func main() {

	r := gin.Default()

	//跨域中间件
	r.Use(hooks.CORSMiddleware())

	//读取JSON文件
	JSON, err := tools.ReadJSON("./config.json")
	if err != nil {
		fmt.Println("读取配置文件失败", err)
	}

	fmt.Println(JSON.ProxyIP, JSON)

	//代理中间件
	hooks.Proxy(r, JSON.ProxyIP)

	fmt.Println("server RUN in listen" + JSON.LocalListenIP)
	fmt.Println("http://localhost:" + JSON.LocalListenIP)

	r.Run(":" + JSON.LocalListenIP)
}
