package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

func UserServiceProducer(message string) {
	config := sarama.NewConfig()

	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	defer producer.Close()

	newMessage := sarama.ProducerMessage{
		Topic: "user-kafka-topic",
		Value: sarama.StringEncoder(message),
	}

	partition, off_set, err := producer.SendMessage(&newMessage)

	if err != nil {
		log.Fatalln("Unable to send message due to: " + err.Error())
	}

	log.Println(partition, off_set)
}
