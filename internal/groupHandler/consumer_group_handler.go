package groupHandler

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/jacksonbarreto/WebGateScanner-DNSSECAnalyzer/pkg/logservice"
	kmodels "github.com/jacksonbarreto/WebGateScanner-kafka/models"
	"github.com/jacksonbarreto/WebGateScanner-stls/config"
	"github.com/jacksonbarreto/WebGateScanner-stls/scanner"
	"sync"
)

type Scanner interface {
}

type GroupHandler struct {
	scanner *scanner.Scanner
	Log     logservice.Logger
}

func NewConsumerGroupHandler(scanner *scanner.Scanner, logService logservice.Logger) *GroupHandler {
	return &GroupHandler{
		scanner: scanner,
		Log:     logService,
	}
}

func NewConsumerGroupHandlerDefault(scanner *scanner.Scanner) *GroupHandler {
	return NewConsumerGroupHandler(scanner, logservice.NewLogService(config.App().Id))
}

func (h *GroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *GroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *GroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var wg sync.WaitGroup
	for message := range claim.Messages() {
		wg.Add(1)
		go func(msg *sarama.ConsumerMessage) {
			defer wg.Done()
			h.Log.Info("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)

			var evalRequest kmodels.EvaluationRequest
			err := json.Unmarshal(msg.Value, &evalRequest)
			if err != nil {
				h.Log.Error("Error unmarshalling message: %v", err)
				return
			}

			err = h.scanner.Scan(evalRequest.URL)
			if err != nil {
				h.Log.Error("Error scanning host '%s': %v", evalRequest.URL, err)
				return
			}
			h.Log.Info("Message processed: value = %s, timestamp = %v, topic = %s", evalRequest.URL, message.Timestamp, message.Topic)
			session.MarkMessage(msg, "")
		}(message)
	}
	wg.Wait()
	return nil
}
