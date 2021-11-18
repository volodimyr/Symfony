package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/volodimyr/Symphony/models"

	"github.com/labstack/echo/v4"
	"github.com/volodimyr/Symphony/domain"
)

const (
	XCarsCount = "X-Cars-Count"
)

type CarsController struct {
	repo domain.CarsRepository

	resolveErr func(error) error
}

func (c CarsController) FindAvailableCars(ctx echo.Context) error {
	start, err := parseDate(ctx, startDateQParam)
	if err != nil {
		return err
	}
	end, err := parseDate(ctx, endDateQParam)
	if err != nil {
		return err
	}
	if err := validateStartEndDate(start, end); err != nil {
		return err
	}

	cars, err := c.repo.GetAvailableCarList(start, end.Add(time.Hour*23+time.Minute*59+time.Second*59))
	if err != nil {
		return err
	}

	response := make([]map[string]interface{}, 0, len(cars))
	for _, car := range cars {
		response = append(response, c.respondCar(car))
	}
	ctx.Response().Header().Set(XCarsCount, strconv.Itoa(len(cars)))

	return ctx.JSON(http.StatusOK, response)
}

func (c CarsController) FindByID(ctx echo.Context) error {
	id, err := getPathParamAsInt(ctx, idPathParam)
	if err != nil {
		return err
	}

	car, err := c.repo.GetByID(id)
	if err != nil {
		return c.resolveErr(err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":    car.ID,
		"model": car.Model,
		"color": car.Color,
	})
}

func (c CarsController) FindAll(ctx echo.Context) error {
	cars, err := c.repo.GetList()
	if err != nil {
		return err
	}

	response := make([]map[string]interface{}, 0, len(cars))
	for _, car := range cars {
		response = append(response, c.respondCar(car))
	}
	ctx.Response().Header().Set(XCarsCount, strconv.Itoa(len(cars)))

	return ctx.JSON(http.StatusOK, response)
}

func (c CarsController) respondCar(car *models.Car) map[string]interface{} {
	return map[string]interface{}{
		"id":    car.ID,
		"model": car.Model,
		"color": car.Color,
	}
}
