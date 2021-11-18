package mock

import (
	"time"

	"github.com/volodimyr/Symphony/models"
)

type CarOrdersRepository struct {
	IsCarAvailableFunc func(carID int, start, end time.Time) error
	GetByIDFunc        func(id int) (*models.CarOrder, error)
	CancelCarOrderFunc func(order *models.CarOrder) error
	OrderCarFunc       func(carID int, start, end time.Time) (models.CarOrder, error)
}

func (c CarOrdersRepository) IsCarAvailable(carID int, start, end time.Time) error {
	return c.IsCarAvailableFunc(carID, start, end)
}

func (c CarOrdersRepository) GetByID(id int) (*models.CarOrder, error) {
	return c.GetByIDFunc(id)
}

func (c CarOrdersRepository) CancelCarOrder(order *models.CarOrder) error {
	return c.CancelCarOrderFunc(order)
}

func (c CarOrdersRepository) OrderCar(carID int, start, end time.Time) (models.CarOrder, error) {
	return c.OrderCarFunc(carID, start, end)
}
