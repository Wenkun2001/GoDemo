package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	topic  = "user_click"
	reader *kafka.Reader
)

// 生产消息
func writeKafka(ctx context.Context) {
	writer := kafka.Writer{
		// 不定长参数，支持传入多个broker的ip:port
		Addr: kafka.TCP("localhost:9092"),
		// 把所有message指定统一的topic。
		// 如果这里不指定统一的Topic，则创建kafka.Message{}时需要分别指定Topic
		Topic: topic,
		// 把所有message的key进行hash，确定partition
		Balancer: &kafka.Hash{},
		// 设定写超时
		WriteTimeout: 1 * time.Second,
		// RequireNone不需要等待ack返回，效率最高，安全性最低；
		// RequireOne只需要确保Leader写入成功就可以发送下一条消息；
		// RequiredAcks需要确保Leader和所有Follower都写入成功才可以发送下一条消息。
		RequiredAcks: kafka.RequireNone,
		// Topic不存在时自动创建。
		// 生产环境中一般设为false，由运维管理员创建Topic并配置partition数目。
		AllowAutoTopicCreation: true,
	}
	// 关闭连接
	defer writer.Close()

	// 允许重试3次
	for i := 0; i < 3; i++ {
		// 批量写入消息，原子操作，要么全写成功，要么全写失败
		if err := writer.WriteMessages(ctx,
			kafka.Message{Key: []byte("1"), Value: []byte("秦")},
			kafka.Message{Key: []byte("2"), Value: []byte("始")},
			kafka.Message{Key: []byte("3"), Value: []byte("皇")},
			// key相同时肯定写入同一个partition
			kafka.Message{Key: []byte("2"), Value: []byte("代")},
			kafka.Message{Key: []byte("3"), Value: []byte("码")},
		); err != nil {
			// 首次写一个新的Topic时，会发生LeaderNotAvailable错误，重试一次
			if err == kafka.LeaderNotAvailable {
				time.Sleep(500 * time.Millisecond)
				continue
			} else {
				fmt.Printf("batch write message failed: %v", err)
			}
		} else {
			//只要成功一次就不再尝试下一次了
			break
		}
	}
}

// 消费消息
func readKafka(ctx context.Context) {
	reader = kafka.NewReader(kafka.ReaderConfig{
		// 支持传入多个broker的ip:port
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
		// 每隔多长时间自动commit一次offset。即一边读一边向kafka上报读到了哪个位置。
		CommitInterval: 1 * time.Second,
		// 一个Group内消费到的消息不会重复
		GroupID: "recommend_biz",
		// 当一个特定的partition没有commited offset时(比如第一次读一个partition，之前没有commit过)，
		// 通过StartOffset指定从第一个还是最后一个位置开始消费。
		// StartOffset的取值要么是FirstOffset要么是LastOffset，
		// LastOffset表示Consumer启动之前生成的老数据不管了。
		// 仅当指定了GroupID时，StartOffset才生效。
		StartOffset: kafka.FirstOffset,
	})
	// 由于下面是死循环，正常情况下readKafka()函数永远不会结束，defer不会执行。
	// 所以需要监听信息2和15，当收到信号时关闭reader。需要把reader设为全局变量
	// defer reader.Close()

	// 消息队列里随时可能有新消息进来，所以这里是死循环，类似于读Channel
	for {
		if message, err := reader.ReadMessage(ctx); err != nil {
			fmt.Printf("read message from kafka failed: %v", err)
			break
		} else {
			offset := message.Offset
			fmt.Printf("topic=%s, partition=%d, offset=%d, key=%s, message content=%s\n", message.Topic, message.Partition, offset, string(message.Key), string(message.Value))
		}
	}
}

// 需要监听信息2和15，当收到信号时关闭reader
func listenSignal() {
	c := make(chan os.Signal, 1)
	// 注册信号2和15
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	// 阻塞，直到信号的到来
	sig := <-c
	fmt.Printf("receive signal %s\n", sig.String())
	if reader != nil {
		reader.Close()
	}
	// 进程退出
	os.Exit(0)
}

func main() {
	ctx := context.Background()
	// writeKafka(ctx)

	go listenSignal()
	readKafka(ctx)
}
