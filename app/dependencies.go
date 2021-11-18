package app

import (
	"github.com/volodimyr/Symphony/app/carorders"
	"github.com/volodimyr/Symphony/app/cars"
	"github.com/volodimyr/Symphony/domain"
)

type Dependencies struct {
	CarsRepo        domain.CarsRepository
	CarOrdersRepo   domain.CarOrdersRepository
	PubSubRepo      domain.PubSubRepository
	CarOrderService domain.CarOrderService
}

func NewDependencies(c Config) (Dependencies, error) {
	rc, err := NewRedisClient(c.Redis)
	if err != nil {
		return Dependencies{}, err
	}

	postgres, err := ConnectPostgresDB(c.Postgres)
	if err != nil {
		return Dependencies{}, err
	}
	if err := MigratePostgresDBSchema(postgres); err != nil {
		return Dependencies{}, err
	}

	d := Dependencies{}
	d.CarsRepo = cars.NewRepository(postgres)
	d.CarOrdersRepo = carorders.NewRepository(postgres)
	d.PubSubRepo = carorders.NewRedisRepository(rc)
	d.CarOrderService = carorders.NewService(d.CarOrdersRepo, d.PubSubRepo)

	return d, nil
}
