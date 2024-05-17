package comsumer

import (
	"log"

	"github.com/Shopify/sarama"
)

func Comsumekafka() {
	
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer([]string{"192.168.89.184:9092"}, config)
	if err != nil {
		log.Fatal(err)
	}
	
	defer consumer.Close()
	
	partitionConsumer, err := consumer.ConsumePartition("killers", 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Consuming partition: %v", err)
	}
	
	defer partitionConsumer.Close()

	
	for {
		select {
		
		case msg := <-partitionConsumer.Messages():
			log.Printf("Received message: %s", string(msg.Value))
		
		case err := <-partitionConsumer.Errors():
			log.Fatalf("Received error: %v", err)
		}
	}
}
