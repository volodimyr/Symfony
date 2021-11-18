//+build unit

package http

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volodimyr/Symphony/domain"
	"github.com/volodimyr/Symphony/domain/mock"
	"github.com/volodimyr/Symphony/models"
)

func TestCarOrdersController_CancelOrder(t *testing.T) {
	const path = "/api/v1/orders/{id}"

	t.Run("invalid id", func(t *testing.T) {
		//GIVEN
		expect, done := initTest(t, emptyRouterInitializer)

		defer done()
		//WHEN
		resp := expect.DELETE(path).
			WithPath("id", "aaa").
			Expect()

		//THEN
		resp.Status(http.StatusUnprocessableEntity).JSON().Equal(map[string]string{
			"message": "id [aaa] has invalid format",
		})
	})

	t.Run("get by id internal server error", func(t *testing.T) {
		//GIVEN
		expect, done := initTest(t, func(router *Router) {
			router.D.CarOrdersRepo = mock.CarOrdersRepository{
				GetByIDFunc: func(id int) (order *models.CarOrder, err error) {
					return nil, fmt.Errorf("went wrong")
				},
			}
		})

		defer done()
		//WHEN
		resp := expect.DELETE(path).
			WithPath("id", 1).
			Expect()

		//THEN
		resp.Status(http.StatusInternalServerError).JSON().Equal(map[string]string{
			"message": "Internal Server Error",
		})
	})

	t.Run("not found", func(t *testing.T) {
		//GIVEN
		expect, done := initTest(t, func(router *Router) {
			router.D.CarOrdersRepo = mock.CarOrdersRepository{
				GetByIDFunc: func(id int) (order *models.CarOrder, err error) {
					return nil, domain.ErrOrdersNotFound
				},
			}
		})

		defer done()
		//WHEN
		resp := expect.DELETE(path).
			WithPath("id", 1).
			Expect()

		//THEN
		resp.Status(http.StatusNotFound).JSON().Equal(map[string]string{
			"message": domain.ErrOrdersNotFound.Error(),
		})
	})

	t.Run("cancel car order internal server error", func(t *testing.T) {
		//GIVEN
		expect, done := initTest(t, func(router *Router) {
			router.D.CarOrdersRepo = mock.CarOrdersRepository{
				GetByIDFunc: func(id int) (order *models.CarOrder, err error) {
					return &models.CarOrder{
						ID:    1,
						CarID: 2,
					}, nil
				},

				CancelCarOrderFunc: func(order *models.CarOrder) error {
					assert.Equal(t, 1, order.ID)
					assert.Equal(t, 2, order.CarID)

					return fmt.Errorf("went wrong")
				},
			}
		})

		defer done()
		//WHEN
		resp := expect.DELETE(path).
			WithPath("id", 1).
			Expect()

		//THEN
		resp.Status(http.StatusInternalServerError).JSON().Equal(map[string]string{
			"message": "Internal Server Error",
		})
	})

}
