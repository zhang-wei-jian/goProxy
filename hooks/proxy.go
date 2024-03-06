package hooks

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"net/url"
)

func Proxy(r *gin.Engine, proxyPath string) {

	// 定义反向代理的目标地址
	//targetURL := "http://localhost:8898"
	//targetURL := "http://127.0.0.1:8898"
	targetURL := proxyPath
	// 创建一个反向代理的 Transport
	fmt.Println("PRxuy")
	target, _ := url.Parse(targetURL)
	proxy := httputil.NewSingleHostReverseProxy(target)

	// 定义路由，将所有请求都代理到目标地址
	r.Any("/*path", func(c *gin.Context) {

		c.Request.Host = target.Host
		// 执行反向代理
		proxy.ServeHTTP(c.Writer, c.Request)

		c.JSON(200, "aaaa")
	})

	//不应该是中间件
}
