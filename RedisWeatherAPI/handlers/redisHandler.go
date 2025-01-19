package handlers

import (
	"context"
	"fmt"
	"log"


	"github.com/redis/go-redis/v9"
)

func InitClient(){
  log.Println("init function is called ")
  ctx := context.Background()
client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Username: "ashfaq",  // Optional (Redis 6+)
        Password: "hellohs72oi",  // Required
    })
pong, err := client.Ping(ctx).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(pong)

  error := client.Set(ctx, "foo", "bar", 0).Err()
  if error != nil {
    log.Println("unable to set value")
    log.Fatalln(error)
  }

val, err := client.Get(ctx, "foo").Result()
if err != nil {
    panic(err)
}
fmt.Println("foo", val)
}
