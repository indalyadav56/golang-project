package config

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {

  // Create client
  rdb := redis.NewClient(&redis.Options{
    Addr: "localhost:6379", 
  })

  // Ping redis
  err := rdb.Ping(context.Background()).Err()
  if err != nil {
    panic(err)
  }

  // Basic SET and GET 
  err = rdb.Set(context.Background(), "key", "value", 0).Err()
  if err != nil {
    panic(err)
  }

  val, err := rdb.Get(context.Background(), "key").Result()
  if err != nil {
    panic(err)
  }
  fmt.Println("key", val)

  // Output: key value
}