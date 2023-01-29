package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0
	// 等待所有follower都回复ack，确保Kafka不会丢消息
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 确认返回，同步生产者需要加这个
	config.Producer.Return.Successes = true
	// 对Key进行Hash，同样的Key每次都落到一个分区，这样消息是有序的
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.ClientID = "go_kafka_client"

	//verbose debugging (comment this line to disabled verbose sarama logging)
	sarama.Logger = log.New(os.Stdout, "[kafka_client] ", log.LstdFlags)

	// 使用同步producer，异步模式下有更高的性能，但是处理更复杂，这里建议先从简单的入手
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	// 关闭消费者，是否资源
	defer func() {
		_ = producer.Close()
	}()
	// 构建生产者失败，抛异常，退出
	if err != nil {
		panic(err.Error())
	}

	msgCount := 4
	// 模拟4个消息
	for i := 0; i < msgCount; i++ {
		rand.Seed(int64(time.Now().Nanosecond()))
		// 构建消息对象，包括主题和消息的key value，value是必须的，key可以不需要
		msg := &sarama.ProducerMessage{
			Topic: "test_go",
			Value: sarama.StringEncoder("hello+" + strconv.Itoa(rand.Intn(100))),
			Key:   sarama.StringEncoder("BBB"),
		}

		t1 := time.Now().Nanosecond()
		// 发送消息，如果成功，返回消息的分区和offset；如果失败，err不为空
		partition, offset, err := producer.SendMessage(msg)
		t2 := time.Now().Nanosecond()

		if err == nil {
			fmt.Println("produce success, partition:", partition, ",offset:", offset, ",cost:", (t2-t1)/(1000*1000), " ms")
		} else {
			fmt.Println(err.Error())
		}

		//client, _ := sarama.NewClient([]string{"localhost:9092"}, config)
		//producer2, _ := sarama.NewAsyncProducerFromClient(client)
		//// 异步发送
		//for i := 0; i <= 100; i++ {
		//	text := fmt.Sprintf("message %08d", i)
		//	producer2.Input() <- &sarama.ProducerMessage{
		//		Topic: "test_go",
		//		Key:   nil,
		//		Value: sarama.StringEncoder(text)}
		//}
	}

}
