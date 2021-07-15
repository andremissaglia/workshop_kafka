package replication

import (
	"context"
	// "fmt"
	// "strconv"

	"github.com/Shopify/sarama"
)

type Worker interface {
	Run(ctx context.Context) error
}

type worker struct {
	storeRatingGateway StoreRatingGateway
	consumer           sarama.ConsumerGroup
	topic              string
}

func NewWorker(
	brokers []string,
	topic string,
	groupID string,
	storeRatingGateway StoreRatingGateway,
) Worker {
	config := sarama.NewConfig()
	config.Version = sarama.MaxVersion
	config.Consumer.Return.Errors = true
	// TODO: Implement

	// go func() {
	// 	for err := range consumer.Errors() {
	// 		fmt.Println("ERROR", err)
	// 	}
	// }()
	return &worker{
		storeRatingGateway: storeRatingGateway,
		topic:              topic,
	}
}

func (w *worker) Run(ctx context.Context) error {
	for {
		err := w.consumer.Consume(ctx, []string{w.topic}, w)
		if err != nil {
			panic(err)
		}
	}
}

func (*worker) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}
func (*worker) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}
func (w *worker) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// TODO: Implement
	return nil
}
