package bootstrap

import (
	"base/src/infra/queue"
	"base/src/present/consumers"
	"go.uber.org/fx"
)

func BuildMessageQueueModules() fx.Option {
	return fx.Options(
		fx.Invoke(queue.NewKafkaProducer),

		fx.Invoke(consumers.NewCDCConsumer),
	)
}
