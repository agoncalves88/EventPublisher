package infrastructure

import (
	"context"

	"EventPublisher.Api/configuration"
	"github.com/segmentio/kafka-go"
)

type KafkaBase struct {
	Configuration configuration.Configuration
}

func (k KafkaBase) Produce(ctx context.Context, message []byte, topic string) error {

	if k.Configuration.BrokerAddress == "" {
		panic("configuration.BrokderAddress can not be nil")
	}
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{k.Configuration.BrokerAddress},
		Topic:   topic,
	})
	err := w.WriteMessages(ctx, kafka.Message{
		Key:   nil,
		Value: message,
	})
	return err
}
