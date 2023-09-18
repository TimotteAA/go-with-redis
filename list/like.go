package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func Like(ctx context.Context, client redis.Cmdable) {
	// 某篇文章的点赞list的key
	key := "post:12345:like"
	for i := 0; i < 100000000; i++ {
		user := fmt.Sprintf("用户:%d", i);
		println("用户 " + user + "点赞了")
		client.RPush(ctx, key, "用户 " + user + "点赞了");
		time.Sleep(time.Millisecond * 100)
	}
}

func GetLike(ctx context.Context, client redis.Cmdable) {
	key := "post:12345:like"
	// 定时5s拉列表前500个
	for {
		time.Sleep(5 * time.Second);
		result, _ := client.LRange(ctx, key, 0, 499).Result()
		for _, val := range result {
			println("val ", val)
		}
		// 删掉前500个
		client.LTrim(ctx, key, 500, -1);
	}
}

func main3() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	go Like(ctx, rdb);
	go GetLike(ctx, rdb);

	select {}
}

