package events

import (
	"encoding/json"
	"wishlist/pkg/models"

	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaPublisher struct {
	Writer *kafka.Writer
}

func NewKafkaPublisher(brokers []string, topic string) *KafkaPublisher {
	return &KafkaPublisher{
		Writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *KafkaPublisher) PublishWishlistItemAdded(ctx context.Context, item *models.WishlistItem) error {
	event := WishlistItemAddedEvent{
		ID:        item.ID,
		UserID:    item.UserID,
		ProductID: item.ProductID,
		AddedAt:   item.AddedAt,
	}

	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(item.ID),
		Value: payload,
	}

	return p.Writer.WriteMessages(ctx, msg)
}
