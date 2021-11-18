package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/volodimyr/Symphony/domain"
)

type CarOrdersController struct {
	repo    domain.CarOrdersRepository
	service domain.CarOrderService

	resolveErr func(error) error
}

func (c CarOrdersController) CancelOrder(ctx echo.Context) error {
	id, err := getPathParamAsInt(ctx, idPathParam)
	if err != nil {
		return err
	}

	order, err := c.repo.GetByID(id)
	if err != nil {
		return c.resolveErr(err)
	}
	if err := c.repo.CancelCarOrder(order); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "ok",
	})
}

func (c CarOrdersController) OrderCar(ctx echo.Context) error {
	carID, err := getPathParamAsInt(ctx, idPathParam)
	if err != nil {
		return err
	}

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

	order, err := c.service.OrderCar(carID, start, end)
	if err != nil {
		return c.resolveErr(err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"order_id":   order.ID,
		"car_id":     order.CarID,
		"start_date": order.StartAt,
		"end_date":   order.EndAt,
	})
}
