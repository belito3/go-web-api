package middleware

import (
	"fmt"
	"github.com/belito3/go-web-api/app/config"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
	"golang.org/x/time/rate"
	"net/http"
	"strconv"
	"time"
)

// App rate limiter implement token bucket algorithm
func ARateLimiterMiddleware(conf config.AppConfiguration, skippers ...SkipperFunc) gin.HandlerFunc {
	cfg := conf.ARateLimiter
	if !cfg.Enable {
		return EmptyMiddleware()
	}

	bucketSize := cfg.Count * 3
	limiter := rate.NewLimiter(rate.Limit(cfg.Count), bucketSize)

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		if limiter.Allow() == false {
			responseError(c, http.StatusTooManyRequests, fmt.Sprintf("App rate limit api exceeded!"))
			c.Abort()
			return
		}

		c.Next()
	}
}


// RateLimiterMiddleware Request frequency limit middleware
// Refs: https://github.com/LyricTian/gin-admin/blob/master/internal/app/middleware/mw_rate_limiter.go
// https://www.alexedwards.net/blog/how-to-rate-limit-http-requests
// TODO: ? sử dung version redis, redis_rate version cao hơn để có context
func CRateLimiterMiddleware(conf config.AppConfiguration, skippers ...SkipperFunc) gin.HandlerFunc {
	cfg := conf.CRateLimiter
	if !cfg.Enable {
		return EmptyMiddleware()
	}

	rc := conf.Redis
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": rc.Addr,
		},
		Password: rc.Password,
		DB:       cfg.RedisDB,
		PoolTimeout: time.Minute, // cancel client that running > timeout
	})

	limiter := redis_rate.NewLimiter(ring)
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}
		ctx := c.Request.Context()
		clientIP := c.ClientIP()
		//userID := ginplus.GetUserID(c)
		if clientIP != "" {
			limit := cfg.Count
			res, err := limiter.Allow(ctx, clientIP, redis_rate.PerMinute(limit))
			//fmt.Println("allowed", res.Allowed, "remaining", res.Remaining)
			if err != nil {
				responseError(c, http.StatusInternalServerError, "Internal Server Error!")
				c.Abort()
				return
			}

			//r, delay, allowed := limiter.Allow(clientIP, limit, time.Second)
			remaining, reset, allowed := res.Remaining, res.ResetAfter, res.Allowed > 0
			h := c.Writer.Header()
			h.Set("X-RateLimit-Limit", strconv.Itoa(limit))
			h.Set("X-RateLimit-Remaining", strconv.Itoa(remaining))
			delaySec := int64(reset / time.Millisecond)
			h.Set("X-RateLimit-Reset", strconv.FormatInt(delaySec, 10))
			if !allowed{
				responseError(c, http.StatusTooManyRequests, fmt.Sprintf("API rate limit exceeded for %v", clientIP))
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

// responseError
func responseError(c *gin.Context, statusCode int, description string){
	c.JSON(statusCode, gin.H{
		"ok": false,
		"error": statusCode,
		"description": description,
	})
}
