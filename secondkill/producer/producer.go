package producer

import (
	"log"
	"time"

	"github.com/Shopify/sarama"
)

func ProduceKafka() {
	// 配置 kafka 生产者
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	// 创建 Kafka 生产者
	producer, err := sarama.NewSyncProducer([]string{"192.168.89.184:9092"}, config)
	if err != nil {
		log.Fatalf("Creating producer: %v", err)
	}
	// 延迟关闭生产者链接
	defer producer.Close()

	// 定义消息 Topic是 go-test, 值为 Hello Kafka
	message := &sarama.ProducerMessage{
		Topic: "killers",
		Value: sarama.StringEncoder("kill success!"),
	}

	// 发送消息
	for i := 0; i < 10; i++ {
		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Fatalf("Sending message: %v", err)
		}
		log.Printf("Message sent to partition %d at offset %d", partition, offset)
		time.Sleep(time.Second)
	}
}
