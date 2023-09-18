package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

const key = "hot_news"

func UpdateHotNews(ctx context.Context, client redis.Cmdable) {
	for {
		for i := 0; i < 10; i++ {
			news := fmt.Sprintf("当前时间：%v，第%d条热点新闻", time.Now(), i)
			// 更新第i条热点新闻
			client.LSet(ctx, key, int64(i), news);
		}
		time.Sleep(time.Second * 1);
	}
}

func GetHotNews(ctx context.Context, client redis.Cmdable) {
	for {
		// 拉取10条热点新闻
		news, _ := client.LRange(ctx, key, 0, -1).Result();
		for _, new := range news {
			fmt.Println("热点新闻 ", new);
		}
		time.Sleep(time.Second * 2);
	}
}

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	go UpdateHotNews(ctx, rdb);
	go GetHotNews(ctx, rdb);

	select {}
}

