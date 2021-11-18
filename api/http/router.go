package http

import (
	"net/http"

	"github.com/volodimyr/Symphony/domain"

	"github.com/labstack/echo/v4"
	"github.com/volodimyr/Symphony/app"
)

type Router struct {
	D app.Dependencies
}

func NewRouter(d app.Dependencies) Router {
	return Router{D: d}
}

func (r *Router) Attach(e *echo.Echo) {
	resolve := r.newErrorResolver().resolve

	cars := CarsController{resolveErr: resolve, repo: r.D.CarsRepo}
	carOrders := CarOrdersController{resolveErr: resolve, repo: r.D.CarOrdersRepo, service: r.D.CarOrderService}

	v1 := e.Group("/api/v1")
	v1.GET("/cars/:id", cars.FindByID)
	v1.GET("/cars", cars.FindAll)
	v1.GET("/available-cars", cars.FindAvailableCars)

	v1.DELETE("/orders/:id", carOrders.CancelOrder)
	v1.POST("/cars/:id/orders", carOrders.OrderCar)
}

func (r *Router) newErrorResolver() errorResolver {
	return errorResolver{errorMap: map[error]httpError{
		// app error : http error
		domain.ErrCarsNotFound:      newHTTPError(http.StatusNotFound, domain.ErrCarsNotFound.Error()),
		domain.ErrOrdersNotFound:    newHTTPError(http.StatusNotFound, domain.ErrOrdersNotFound.Error()),
		domain.ErrCarIsNotAvailable: newHTTPError(http.StatusUnprocessableEntity, domain.ErrCarIsNotAvailable.Error()),
	}}
}

type errorResolver struct {
	errorMap map[error]httpError
}

func (e errorResolver) resolve(err error) error {
	if result, ok := e.errorMap[err]; ok {
		return result
	}

	return err
}

type httpError struct {
	status int
	msg    string
}

func newHTTPError(code int, message string) httpError {
	return httpError{status: code, msg: message}
}

func (h httpError) Error() string {
	return h.msg
}

func (h httpError) Code() int {
	return h.status
}

func (h httpError) Message() string {
	return h.msg
}
