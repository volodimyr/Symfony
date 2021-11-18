//+build integration

package http

import (
	"net/http"
	"testing"

	"github.com/volodimyr/Symphony/app/carorders"
	"github.com/volodimyr/Symphony/domain"
)

const carOrderDataRes = "carorder_resources"

func TestIntegrationCarOrdersController_CancelOrder(t *testing.T) {
	const path = "/api/v1/orders/{id}"

	t.Run("order not found", func(t *testing.T) {
		//GIVEN
		migrateEmptyTestData(t)
		expect, done := initTest(t, func(router *Router) {
			router.D.CarOrdersRepo = carorders.NewRepository(postgresDBTest)
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

	t.Run("successful cancellation", func(t *testing.T) {
		//GIVEN
		migrateTestData(t, carOrderDataRes+"/carorder_cancel")
		expect, done := initTest(t, func(router *Router) {
			router.D.CarOrdersRepo = carorders.NewRepository(postgresDBTest)
		})
		defer done()

		//WHEN
		resp := expect.DELETE(path).
			WithPath("id", 555).
			Expect()

		//THEN
		resp.Status(http.StatusOK).JSON().Equal(map[string]string{
			"message": "ok",
		})
	})
}
