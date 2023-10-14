package akafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

func Consume(topics []string, servers string, msgChan chan *kafka.Message){
  kafkaConsumer, err := kafka.NewConsumer(&kafka.configMap{
  "bootstrap.servers": servers
  "group.id": "hexagonal-akafka-golang-crud-product" 
  "auto.offset.reset": "earliest" 
  }
  if err != nil{
    panic(err)
  }

  kafkaConsumer.SubscribeTopics(topics, nil)
  for{
    msg, err := kafkaConsumer.ReadMessage(-1)
    if err == nil {
       msgChan <- msg
    }
  }

}
