package ratelimit

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// RateLimiter 基于 Redis 的限流器
type RateLimiter struct {
	rdb *redis.Client
}

// NewRateLimiter 创建限流器
func NewRateLimiter(rdb *redis.Client) *RateLimiter {
	return &RateLimiter{rdb: rdb}
}

// Lua 脚本实现原子限流（固定窗口算法）
// KEYS[1]: 限流标识符（如 ip + route）
// ARGV[1]: 限制请求次数
// ARGV[2]: 时间窗口大小（秒）
var limitLuaScript = redis.NewScript(`
local key = KEYS[1]
local limit = tonumber(ARGV[1])
local window = tonumber(ARGV[2])

local current = redis.call('INCR', key)
if current == 1 then
    redis.call('EXPIRE', key, window)
end

if current > limit then
    return 0
end
return 1
`)

// Allow 判断是否允许请求
func (l *RateLimiter) Allow(ctx context.Context, key string, limit int, window time.Duration) (bool, error) {
	// 将 duration 转换为整数秒，向上取整确保不小于 1 秒
	seconds := int(window.Seconds())
	if seconds <= 0 {
		seconds = 1
	}

	res, err := limitLuaScript.Run(ctx, l.rdb, []string{key}, limit, seconds).Int()
	if err != nil {
		return false, err
	}

	return res == 1, nil
}
