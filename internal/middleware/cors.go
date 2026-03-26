package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, X-Request-ID, Accept-Language")
		c.Header("Access-Control-Expose-Headers", "X-Request-ID")
		// Max-Age 单位是秒，浏览器在此时间内缓存预检结果，不再重复发 OPTIONS
		c.Header("Access-Control-Max-Age", "43200") // 12 小时

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
