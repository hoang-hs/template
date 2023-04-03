package consumers

import (
	"base/src/common"
	"base/src/common/configs"
	"base/src/common/log"
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/fx"
	"runtime/debug"
	"time"
)

const (
	timeoutKafka = 2 * time.Second
)

type CDCConsumer struct {
	consumer *kafka.Consumer
}

func NewCDCConsumer(lc fx.Lifecycle) {
	cf := configs.Get().Kafka
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cf.Host,
		//Todo add group
		"group.id":           "group 1",
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
	})
	if err != nil {
		log.Fatal("Failed to create new consumer, err:[%s]", err.Error())
	}
	//Todo add topic
	err = c.SubscribeTopics([]string{""}, nil)
	if err != nil {
		log.Fatal("Failed to subscribe topic, topic:[%s]", []string{})
	}
	quit := make(chan bool)
	cdcConsumer := &CDCConsumer{
		consumer: c,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.GetLogger().GetZap().Info("OnStart CDC consumer")
			go cdcConsumer.Run(quit)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.GetLogger().GetZap().Info("OnStop CDC consumer")
			quit <- true
			return cdcConsumer.consumer.Close()
		},
	})
}

func (c *CDCConsumer) Run(quit chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			func() {
				ctx := common.CreateNewCtx()
				defer recovery(ctx)
				msg, err := c.consumer.ReadMessage(timeoutKafka)
				if err != nil {
					if err.(kafka.Error).IsTimeout() == false {
						log.Error(ctx, "read message kafka error, err:[%s]", err.Error())
					}
					return
				}
				log.Info(ctx, "Receive message, topic: [%s], offset: [%s], partition: [%d]",
					*msg.TopicPartition.Topic, msg.TopicPartition.Offset.String(), msg.TopicPartition.Partition)
				switch *msg.TopicPartition.Topic {

				default:
					log.Warn(ctx, "Receive message not handle: ", *msg.TopicPartition.Topic, msg.TopicPartition.Offset.String(), err)
				}
			}()

		}
	}

}

func recovery(ctx context.Context) {
	if err := recover(); err != nil {
		log.Error(ctx, "[Recovery from panic]\ntime: [%v]\nerror: [%v]\nstack: [%v]\n",
			time.Now(), err, string(debug.Stack()))
	}
}
