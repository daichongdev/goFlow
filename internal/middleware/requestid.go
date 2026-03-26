package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"goflow/internal/pkg/logger"
)

const RequestIDHeader = "X-Request-ID"

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader(RequestIDHeader)
		if reqID == "" {
			reqID = uuid.New().String()
		}
		c.Set(string(logger.RequestIDKey), reqID)
		c.Header(RequestIDHeader, reqID)

		// 写入 request context，供 logger.WithCtx 使用
		ctx := context.WithValue(c.Request.Context(), logger.RequestIDKey, reqID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
