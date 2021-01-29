package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

const (
	topic          = "event-create-account"
	broker1Address = "localhost:9092"
)

func Produce(ctx context.Context, message []byte) {

	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
	})

	err := w.WriteMessages(ctx, kafka.Message{
		Key: nil,
		// create an arbitrary message payload for the value
		Value: message,
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}

}
