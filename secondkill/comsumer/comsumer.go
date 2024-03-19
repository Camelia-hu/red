package comsumer

import (
	"log"

	"github.com/Shopify/sarama"
)

func Comsumekafka() {
	// 配置 Kafka 消费者
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// 创建 Kafka 消费者
	consumer, err := sarama.NewConsumer([]string{"192.168.89.184:9092"}, config)
	if err != nil {
		log.Fatal(err)
	}
	// 延迟关闭消费者链接
	defer consumer.Close()
	//订阅主题，获取分区 partition
	partitionConsumer, err := consumer.ConsumePartition("killers", 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Consuming partition: %v", err)
	}
	// 延迟关闭分区链接
	defer partitionConsumer.Close()

	// 消费消息
	for {
		select {
		// 从 分区 通道中获取信息
		case msg := <-partitionConsumer.Messages():
			log.Printf("Received message: %s", string(msg.Value))
		// 如果从通道中获取消息失败
		case err := <-partitionConsumer.Errors():
			log.Fatalf("Received error: %v", err)
		}
	}
}
