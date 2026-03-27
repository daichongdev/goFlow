package middleware

import (
	"fmt"
	"time"

	"goflow/internal/pkg/errcode"
	"goflow/internal/pkg/ratelimit"
	"goflow/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

// RateLimit 基于 IP 和路由的限流中间件
// limit: 窗口内允许的最大请求数
// window: 时间窗口大小
func RateLimit(limiter *ratelimit.RateLimiter, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		// 获取注册时的完整路径，如 /app/v1/products/:id
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}

		// 构造 Redis Key: rl:ip:path
		key := fmt.Sprintf("rl:%s:%s", ip, path)

		allowed, err := limiter.Allow(c.Request.Context(), key, limit, window)
		if err != nil {
			// Redis 故障时降级：记录日志并放行，不影响核心业务可用性
			// 也可以选择 c.Abort()，取决于业务对限流的强制程度
			c.Next()
			return
		}

		if !allowed {
			// 返回 429 Too Many Requests
			response.Error(c, errcode.ErrTooManyRequests())
			c.Abort()
			return
		}

		c.Next()
	}
}
