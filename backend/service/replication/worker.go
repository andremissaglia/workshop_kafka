package replication

import (
	"context"
	"fmt"
	"strconv"

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
	consumer, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		panic(err)
	}
	go func() {
		for err := range consumer.Errors() {
			fmt.Println("ERROR", err)
		}
	}()
	return &worker{
		storeRatingGateway: storeRatingGateway,
		consumer:           consumer,
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
	for msg := range claim.Messages() {
		key := string(msg.Key)
		value := string(msg.Value)
		catID, err := strconv.Atoi(key)
		if err != nil {
			return err
		}

		rating, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return err
		}

		err = w.storeRatingGateway.Store(context.Background(), catID, float32(rating))
		if err != nil {
			return err
		}
		sess.MarkMessage(msg, "")
	}
	return nil
}
