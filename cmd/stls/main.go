package main

import (
	"github.com/jacksonbarreto/stls/config"
	"github.com/jacksonbarreto/stls/internal/consumer"
	"github.com/jacksonbarreto/stls/scanner"
)

const configFilePath = ""

func main() {
	config.InitConfig(configFilePath)
	scan := scanner.NewScanner()
	handler := consumer.NewConsumerGroupHandlerDefault(scan)
	kafkaConsumer, consumerErr := consumer.NewConsumerDefault(handler)
	if consumerErr != nil {
		panic(consumerErr)
	}

	consumeErr := kafkaConsumer.Consume()
	if consumeErr != nil {
		panic(consumeErr)
	}
}
