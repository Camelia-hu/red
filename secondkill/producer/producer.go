package producer

import (
	"log"
	"time"

	"github.com/Shopify/sarama"
)

func ProduceKafka() {
	
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	
	producer, err := sarama.NewSyncProducer([]string{"192.168.89.184:9092"}, config)
	if err != nil {
		log.Fatalf("Creating producer: %v", err)
	}
	
	defer producer.Close()

	
	message := &sarama.ProducerMessage{
		Topic: "killers",
		Value: sarama.StringEncoder("kill success!"),
	}

	
	for i := 0; i < 10; i++ {
		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Fatalf("Sending message: %v", err)
		}
		log.Printf("Message sent to partition %d at offset %d", partition, offset)
		time.Sleep(time.Second)
	}
}
