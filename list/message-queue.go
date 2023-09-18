package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

func MQProducer(ctx context.Context, client redis.Cmdable) {
	for {
		// 每隔一秒消费数据
		time.Sleep(time.Second * 5);
		num := rand.Intn(1000)
		_ = client.LPush(ctx, "mq-test", num).Err()
		len, _ := client.LLen(ctx, "mq-test").Result()
		fmt.Printf("往队列中写入数据: %d，队列长度：%d\n", num, len)
	}
}

func MQConsumer(ctx context.Context, client redis.Cmdable) {
	for {
		// 右侧取数据用阻塞操作
		result, _ := client.BRPop(ctx, time.Second * 5, "mq-test").Result()
		fmt.Printf("消费到的数据：%v\n ", result)
	}
}

func main2() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	go MQProducer(ctx, rdb)
	go MQConsumer(ctx, rdb)

	// 使用一个信号来阻止主函数退出
	select {}
}