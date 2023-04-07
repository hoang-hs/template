package queue

import (
	"base/src/common/configs"
	"base/src/common/log"
	"base/src/core/message_queue"
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/fx"
)

const (
	timeFlush = 15 * 1000
)

type KafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer(lc fx.Lifecycle) {
	cf := configs.Get().Kafka
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cf.Host,
	})
	if err != nil {
		log.Fatal("Failed to create kafka producer, err:[%s]", err.Error())
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			p.Flush(timeFlush)
			p.Close()
			return nil
		},
	})
	globalProducer = &KafkaProducer{
		producer: p,
	}
}

func (k *KafkaProducer) Produce(ctx context.Context, msg message_queue.AbstractMessageQueue) {
	if msg == nil {
		log.Error(ctx, "msg null")
		return
	}
	value, err := msg.Payload()
	if err != nil {
		log.Error(ctx, "err:[%s]", err.Error())
		return
	}
	err = k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     msg.Topic(),
			Partition: kafka.PartitionAny,
		},
		Value: value,
	}, nil)
	if err != nil {
		log.Error(ctx, "produce msg queue error, topic:[%s], payload:[%s], err:[%s]",
			*msg.Topic(), string(value), err.Error())
	} else {
		log.Info(ctx, "Produce msg success, topic:[%s]", *msg.Topic())
	}
}
