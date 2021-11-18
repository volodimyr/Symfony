package carorders

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/volodimyr/Symphony/models"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) RedisRepository {
	return RedisRepository{
		client: client,
	}
}

func (r RedisRepository) PublishOrder(order models.CarOrder) error {
	const orderTopicName = "new_orders"

	return r.client.Publish(context.Background(), orderTopicName, NewOrderToPublish(order)).Err()
}

type OrderToPublish struct {
	ID    int       `json:"id"`
	CarID int       `json:"car_id"`
	Start time.Time `json:"start_at"`
	End   time.Time `json:"end_at"`
}

func NewOrderToPublish(order models.CarOrder) *OrderToPublish {
	return &OrderToPublish{
		ID:    order.ID,
		CarID: order.CarID,
		Start: order.StartAt,
		End:   order.EndAt,
	}
}

func (o *OrderToPublish) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}

func (o *OrderToPublish) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &o); err != nil {
		return err
	}

	return nil
}
