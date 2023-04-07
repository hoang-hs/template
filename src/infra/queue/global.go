package queue

import (
	"base/src/core/message_queue"
	"context"
)

var globalProducer *KafkaProducer

func Produce(ctx context.Context, msg message_queue.AbstractMessageQueue) {
	globalProducer.Produce(ctx, msg)
}
