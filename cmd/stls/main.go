package main

import (
	"context"
	"github.com/jacksonbarreto/WebGateScanner-kafka/consumer"
	"github.com/jacksonbarreto/WebGateScanner-stls/config"
	"github.com/jacksonbarreto/WebGateScanner-stls/internal/groupHandler"
	"github.com/jacksonbarreto/WebGateScanner-stls/scanner"
)

const configFilePath = ""

func main() {
	config.InitConfig(configFilePath)
	scan := scanner.NewScanner()
	handler := groupHandler.NewConsumerGroupHandlerDefault(scan)
	kafkaConfig := config.Kafka()
	kafkaConsumer, consumerErr := consumer.NewConsumer(kafkaConfig.Brokers, kafkaConfig.GroupID,
		kafkaConfig.TopicsConsumer, handler, context.Background())
	if consumerErr != nil {
		panic(consumerErr)
	}

	consumeErr := kafkaConsumer.Consume()
	if consumeErr != nil {
		panic(consumeErr)
	}

}
