package implcat

import (
	"context"
	// "strconv"

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
	return &voteKafkaAdapter{
		topic: topic,
	}
}

func (a *voteKafkaAdapter) Vote(ctx context.Context, catID int, vote int) error {
	// TODO: Implement
	return nil
}
