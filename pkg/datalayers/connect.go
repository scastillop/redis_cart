package datalayers

import (
	"fmt"
	"os"
	"strconv"
	"github.com/go-redis/redis"
)

// ConnectDB - Create db connection
func ConnectRedis() {

	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
		Password: "",
		DB: 0,
	})
}
