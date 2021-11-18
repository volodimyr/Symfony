package carorders

import (
	"time"

	"github.com/rs/zerolog/log"
	"github.com/volodimyr/Symphony/domain"
	"github.com/volodimyr/Symphony/models"
)

type Service struct {
	repo       domain.CarOrdersRepository
	pubSubRepo domain.PubSubRepository
}

func NewService(repo domain.CarOrdersRepository, psRepo domain.PubSubRepository) Service {
	return Service{
		repo:       repo,
		pubSubRepo: psRepo,
	}
}

func (s Service) OrderCar(carID int, start, end time.Time) (models.CarOrder, error) {
	end = end.Add(time.Hour*23 + time.Minute*59 + time.Second*59)

	switch err := s.repo.IsCarAvailable(carID, start, end); err {
	case nil:
		// car is available
	default:
		return models.CarOrder{}, err
	}

	order, err := s.repo.OrderCar(carID, start, end)
	if err != nil {
		return models.CarOrder{}, err
	}

	if err := s.pubSubRepo.PublishOrder(order); err != nil {
		log.Err(err).Int("order id", order.ID).Msg("couldn't publish order")
	} else {
		log.Info().Msg("new order was successfully published to the pub sub broker")
	}

	return order, nil
}
