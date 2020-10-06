package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_ = rdb.FlushDB(ctx).Err()

	limiter := redis_rate.NewLimiter(rdb)
	res, err := limiter.Allow(ctx, "project:123", redis_rate.PerMinute(10))
	if err != nil {
		panic(err)
	}
	//
	res, err = limiter.Allow(ctx, "project:123", redis_rate.PerMinute(10))
	res, err = limiter.Allow(ctx, "project:123", redis_rate.PerMinute(10))
	res, err = limiter.Allow(ctx, "project:123", redis_rate.PerMinute(10))
	res, err = limiter.Allow(ctx, "project:123", redis_rate.PerMinute(10))
	fmt.Println("allowed: ", res.Allowed, "remaining: ", res.Remaining, "retry_after: ", res.RetryAfter, "reset_after: ", res.ResetAfter)
	// Output: allowed 1 remaining 9

	time.Sleep(3 * time.Second)

	res, err = limiter.Allow(ctx, "project:123", redis_rate.PerMinute(10))
	res, err = limiter.Allow(ctx, "project:123", redis_rate.PerMinute(10))
	res, err = limiter.Allow(ctx, "project:123", redis_rate.PerMinute(10))
	res, err = limiter.Allow(ctx, "project:123", redis_rate.PerMinute(10))
	res, err = limiter.Allow(ctx, "project:123", redis_rate.PerMinute(10))
	res, err = limiter.Allow(ctx, "project:123", redis_rate.PerMinute(10))
	res, err = limiter.Allow(ctx, "project:123", redis_rate.PerMinute(10))
	res, err = limiter.Allow(ctx, "project:123", redis_rate.PerMinute(10))
	res, err = limiter.Allow(ctx, "project:123", redis_rate.PerMinute(10))
	fmt.Println("allowed: ", res.Allowed, "remaining: ", res.Remaining, "retry_after: ", res.RetryAfter, "reset_after: ", res.ResetAfter)
}