package implcat

import (
	"context"
	"strconv"

	"github.com/andremissaglia/workshop_kafka/backend/service/cat"

	"github.com/Shopify/sarama"
)

type voteKafkaAdapter struct {
	kafkaProducer sarama.SyncProducer
	topic         string
}

func NewVoteKafkaAdapter(brokers []string, topic string) cat.VoteGateway {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Version = sarama.MaxVersion
	producer, err := sarama.NewSyncProducer(
		brokers,
		config,
	)
	if err != nil {
		panic(err)
	}
	return &voteKafkaAdapter{
		kafkaProducer: producer,
		topic:         topic,
	}
}

func (a *voteKafkaAdapter) Vote(ctx context.Context, catID int, vote int) error {
	_, _, err := a.kafkaProducer.SendMessage(&sarama.ProducerMessage{
		Topic: a.topic,
		Key:   sarama.StringEncoder(strconv.Itoa(catID)),
		Value: sarama.StringEncoder(strconv.Itoa(vote)),
	})

	return err
}
