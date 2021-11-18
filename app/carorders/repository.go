package carorders

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

func (r Repository) GetByID(id int) (*models.CarOrder, error) {
	order, err := models.CarOrders(qm.Where("deleted_at IS NULL AND id = ?", id)).One(context.Background(), r.db)

	switch err {
	case sql.ErrNoRows:
		return nil, domain.ErrOrdersNotFound
	case nil:
	// do nothing
	default:
		return nil, fmt.Errorf("failed to get card order by id=%d due to %v", id, err)
	}

	return order, nil
}

func (r Repository) CancelCarOrder(order *models.CarOrder) error {
	order.DeletedAt = null.NewTime(time.Now(), true)
	if _, err := order.Update(context.Background(), r.db, boil.Whitelist(models.CarOrderColumns.DeletedAt)); err != nil {
		return fmt.Errorf("failed to cancel car ordaer by id = %d due to %v", order.ID, err)
	}

	return nil
}

func (r Repository) IsCarAvailable(carID int, start, end time.Time) error {
	count, err := models.CarOrders(
		qm.Where("deleted_at IS NULL"),
		qm.And("car_id = ?", carID),
		qm.And("? BETWEEN start_at AND end_at", start),
		qm.Or("start_at BETWEEN ? AND ?", start, end),
	).Count(context.Background(), r.db)

	if err != nil {
		return fmt.Errorf("failed to check whether car is available")
	}
	if count == 0 {
		return nil
	}

	return domain.ErrCarIsNotAvailable
}

func (r Repository) OrderCar(carID int, start, end time.Time) (models.CarOrder, error) {
	var order models.CarOrder
	order.CarID = carID
	order.StartAt = start
	order.EndAt = end

	if err := order.Insert(context.Background(), r.db,
		boil.Whitelist(models.CarOrderColumns.CarID, models.CarOrderColumns.StartAt, models.CarOrderColumns.EndAt)); err != nil {
		return order, fmt.Errorf("failed to book a car = %d ue to %v", carID, err)
	}

	return order, nil
}
