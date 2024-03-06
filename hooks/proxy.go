package hooks

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Proxy(r *gin.Engine, proxyPath string) {

	// 定义反向代理的目标地址

	//targetURL := proxyPath
	//// 创建一个反向代理的 Transport
	//fmt.Println("PRxuy")
	//target, _ := url.Parse(targetURL)
	//proxy := httputil.NewSingleHostReverseProxy(target)

	target, _ := url.Parse(proxyPath)
	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.ModifyResponse = func(resp *http.Response) error {
		// 删除原有的 Access-Control-Allow-Origin 头首先删除了响应头中的 Access-Control-Allow-Origin 字段。这是因为目标服务器可能已经设置了这个字段，但我们希望使用自己的值，所以我们需要先删除它。
		resp.Header.Del("Access-Control-Allow-Origin")

		// 添加新的 Access-Control-Allow-Origin 头
		//resp.Header.Set("Access-Control-Allow-Origin", "*")

		return nil
	}

	// 定义路由，将所有请求都代理到目标地址
	r.Any("/*path", func(c *gin.Context) {

		// Modify request URL
		c.Request.URL.Scheme = "http"
		c.Request.URL.Host = target.Host
		//c.Request.Host = target.Host
		// 执行反向代理
		proxy.ServeHTTP(c.Writer, c.Request)

	})

	//不应该是中间件
}
