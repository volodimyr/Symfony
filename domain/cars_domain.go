package domain

import (
	"time"

	"github.com/volodimyr/Symphony/models"
)

type CarsRepository interface {
	GetList() (models.CarSlice, error)
	GetByID(id int) (*models.Car, error)
	GetAvailableCarList(start, end time.Time) (models.CarSlice, error)
}
