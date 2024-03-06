package hooks

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 中间件

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Request.Header.Set("X-Forwarded-For", "127.0.0.1")

		// Set X-Forwarded-For header
		c.Request.Header.Set("X-Forwarded-For", "127.0.0.1")
		c.Request.Header.Set("Origin", "http://localhost")
		c.Request.Header.Set("Referer", "http://localhost")

		// 设置允许跨域请求的来源
		//origin := c.Request.Header.Get("Origin")
		//c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		//c.Writer.Header().Set("Access-Control-Allow-Origin", "http://60.204.170.225:8503")
		// 设置允许跨域请求的来源为任何域
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//c.Writer.Header().Del("Access-Control-Allow-Origin")

		// 其他 CORS 头设置
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			c.Next()
		}
	}
}
