package main

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

var (
	ctx     = context.Background()
	rdb     *redis.Client
	limiter *redis_rate.Limiter
)

func main() {
	initRedis("localhost:6379")
	r := gin.Default()
	r.GET("/fib/:n", fibHandler)
	r.Run(":8080")
}

func initRedis(addr string) {
	rdb = redis.NewClient(&redis.Options{Addr: addr})
	limiter = redis_rate.NewLimiter(rdb)
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(err)
	}
}

func fibHandler(c *gin.Context) {
	ip := c.ClientIP()
	res, err := limiter.Allow(ctx, "rate:"+ip, redis_rate.PerSecond(10))

	if err != nil || res.Allowed == 0 {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many request"})
		return
	}

	nSt := c.Param("n")

	value, err := rdb.Get(ctx, "fib-"+nSt).Result()

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"fib": value, "from": "cache"})
		return
	}

	n, err := strconv.ParseUint(nSt, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not number"})
		return
	}

	resp := fib(n)

	if err := rdb.Set(ctx, "fib-"+nSt, resp, time.Hour).Err(); err == nil {
		c.JSON(http.StatusOK, gin.H{"fib": resp, "from": "computed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"fib": resp, "from": "computed"})
}

func fib(n uint64) string {
	if n <= 1 {
		return "1"
	}
	a, b := uint64(1), uint64(1)
	for i := uint64(2); i < n; i++ {
		a, b = b, a+b
	}
	return strconv.FormatUint(b, 10)
}
