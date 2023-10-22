package kafka

import (
	"fmt"
	"os"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
	}

	p, err := ckafka.NewProducer(configMap)

	if err != nil {
		panic(err)
	}

	return p
}

func Publish(msg string, topic string, producer *ckafka.Producer, deliveryChan chan ckafka.Event) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
	}

	err := producer.Produce(message, deliveryChan)

	if err != nil {
		return err
	}

	return nil
}

/*
	Essa função é responsável por ficar escutando o canal de eventos do kafka e imprimir na tela
	quando a mensagem foi entregue ou não. 
	Ela fica rodando em background, por isso é uma goroutine.
	Para fazer com que as funçoes se comuniquem, será criada uma instancia de DeliveryReport na CLI localizada no pasta cmd/kafka.go
*/
func DeliveryReport(deliveryChan chan ckafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *ckafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Delivery failed:", ev.TopicPartition)
			} else {
				fmt.Println("Delivered message to:", ev.TopicPartition)
			}
		}
	}
}