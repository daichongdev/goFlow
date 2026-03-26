package mq

import (
	"strings"

	"goflow/internal/pkg/logger"

	"github.com/ThreeDotsLabs/watermill"
)

// zapLogger 适配 Watermill 的 LoggerAdapter 接口，转接到项目 zap logger
type zapLogger struct{}

func newLogger() watermill.LoggerAdapter {
	return &zapLogger{}
}

func (z *zapLogger) Error(msg string, err error, fields watermill.LogFields) {
	// 多个 subscriber 共享同一 Redis client，关闭时后续 subscriber 会收到
	// "redis: client is closed" 错误，属于预期行为，降级为 debug 日志
	if err != nil && strings.Contains(err.Error(), "client is closed") {
		logger.Log.Debugw(msg, append(fieldsToArgs(fields), "error", err)...)
		return
	}
	args := fieldsToArgs(fields)
	if err != nil {
		args = append(args, "error", err)
	}
	logger.Log.Errorw(msg, args...)
}

func (z *zapLogger) Info(msg string, fields watermill.LogFields) {
	logger.Log.Infow(msg, fieldsToArgs(fields)...)
}

func (z *zapLogger) Debug(msg string, fields watermill.LogFields) {
	logger.Log.Debugw(msg, fieldsToArgs(fields)...)
}

func (z *zapLogger) Trace(msg string, fields watermill.LogFields) {
	logger.Log.Debugw("[trace] "+msg, fieldsToArgs(fields)...)
}

func (z *zapLogger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	return z
}

func fieldsToArgs(fields watermill.LogFields) []interface{} {
	args := make([]interface{}, 0, len(fields)*2)
	for k, v := range fields {
		args = append(args, k, v)
	}
	return args
}
