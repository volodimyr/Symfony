//+build unit integration

package http

import (
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/volodimyr/Symphony/app"
)

func emptyRouterInitializer(_ *Router) {
	//no dependencies needed
}

func initTest(t *testing.T, routerInitializer func(*Router)) (*httpexpect.Expect, func()) {
	r := Router{}
	routerInitializer(&r)

	e := app.NewEchoServer()
	r.Attach(e)

	testServer := httptest.NewServer(e)

	expect := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  testServer.URL,
		Reporter: httpexpect.NewAssertReporter(t),
	})

	return expect, func() {
		testServer.Close()
	}
}
