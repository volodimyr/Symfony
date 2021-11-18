package domain

import (
	"time"

	"github.com/volodimyr/Symphony/models"
)

type CarOrderService interface {
	OrderCar(carID int, start, end time.Time) (models.CarOrder, error)
}

type CarOrdersRepository interface {
	IsCarAvailable(carID int, start, end time.Time) error
	GetByID(id int) (*models.CarOrder, error)
	CancelCarOrder(order *models.CarOrder) error
	OrderCar(carID int, start, end time.Time) (models.CarOrder, error)
}

type PubSubRepository interface {
	PublishOrder(order models.CarOrder) error
}
