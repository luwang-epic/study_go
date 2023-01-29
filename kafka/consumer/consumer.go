package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
)

type MyConsumerGroup struct{
}

// Setup 会在新会话开始前运行
func (cg *MyConsumerGroup) Setup(_ sarama.ConsumerGroupSession) error {
	fmt.Println("invoke setup...")
	return nil
}

// Cleanup 所有订阅者协程退出后运行
func (cg *MyConsumerGroup) Cleanup(_ sarama.ConsumerGroupSession) error {
	fmt.Println("invoke cleanup...")
	return nil
}

// ConsumeClaim 订阅者消费消息后调用
func (cg *MyConsumerGroup) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	fmt.Println("invoke ConsumeClaim...")
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d  value:%s\n",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Value))

		// 标记，sarama会自动进行提交，默认间隔1秒
		sess.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	consumerConfig := sarama.NewConfig()
	consumerConfig.Version = sarama.V2_8_0_0 // specify appropriate version
	consumerConfig.Consumer.Return.Errors = false
	//consumerConfig.Consumer.Offsets.AutoCommit.Enable = true      // 禁用自动提交，改为手动
	//consumerConfig.Consumer.Offsets.AutoCommit.Interval = time.Second * 1 // 测试3秒自动提交
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetNewest

	cGroup, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "testgroup", consumerConfig)
	if err != nil {
		panic(err)
	}

	cg := &MyConsumerGroup{}
	for {
		err := cGroup.Consume(context.Background(), []string{"test_go"}, cg)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}

	_ = cGroup.Close()
}
