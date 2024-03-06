package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhang-wei-jian/goProxy/hooks"
	"github.com/zhang-wei-jian/goProxy/tools"
)

func main() {

	r := gin.Default()

	r.Use(hooks.CORSMiddleware())

	//读取JSON文件
	JSON, err := tools.ReadJSON("./config.json")
	if err != nil {
		fmt.Println("读取配置文件失败", err)
	}

	fmt.Println(JSON.ProxyIP)

	hooks.Proxy(r, JSON.ProxyIP)
	//r.GET("", func(c *gin.Context) {
	//	c.JSON(200, "aaa")
	//})

	fmt.Println("运行在了 http://localhost:" + JSON.LocalListenIP)
	r.Run(":" + JSON.LocalListenIP)
}
