package cars

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/volatiletech/sqlboiler/v4/queries"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volodimyr/Symphony/domain"
	"github.com/volodimyr/Symphony/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db: db}
}

func (r Repository) GetList() (models.CarSlice, error) {
	cars, err := models.Cars(qm.Where("deleted_at IS NULL")).All(context.Background(), r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get list of cars due to %v", err)
	}

	return cars, nil
}

func (r Repository) GetByID(id int) (*models.Car, error) {
	car, err := models.Cars(qm.Where("deleted_at IS NULL AND id = ?", id)).One(context.Background(), r.db)

	switch err {
	case sql.ErrNoRows:
		return nil, domain.ErrCarsNotFound
	case nil:
	// do nothing
	default:
		return nil, fmt.Errorf("failed to get car by id=%d due to %v", id, err)
	}

	return car, nil
}

func (r Repository) GetAvailableCarList(start, end time.Time) (models.CarSlice, error) {
	cars := models.CarSlice{}
	err := queries.Raw(`
		SELECT cars.id, cars.model, cars.color FROM cars WHERE cars.deleted_at IS NULL AND cars.id NOT IN(
    		SELECT car_id FROM car_orders
        		WHERE car_orders.deleted_at IS NULL 
        		AND ($1 BETWEEN car_orders.start_at AND car_orders.end_at OR (car_orders.start_at BETWEEN $1 AND $2)) GROUP BY car_id
    		);
	`, start, end).Bind(context.Background(), r.db, &cars)
	if err != nil {
		return nil, fmt.Errorf("failed to get list of available cars due to %v", err)
	}

	return cars, nil
}
