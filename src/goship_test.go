package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestMain(t *testing.T) {
	Inst.StartInstance()
	Inst.URLLoader()

	app := Inst.App

	e := httptest.New(t, app)

	e.GET("/").Expect().Status(httptest.StatusOK)
	e.GET("/anchored/28").Expect().Status(httptest.StatusOK)
	e.GET(("/arrivalPrevisions/28")).Expect().Status(httptest.StatusOK)
	e.GET(("/departurePrevisions/28")).Expect().Status(httptest.StatusOK)
	e.GET(("/shiftingPrevisions/28")).Expect().Status(httptest.StatusOK)
	e.GET(("/arrivalsToday/28")).Expect().Status(httptest.StatusOK)
	e.GET(("/departuresToday/28")).Expect().Status(httptest.StatusOK)

	e.GET(("/arrivalsRegister/28/2019-01-01 00:00/2019-01-01 00:00/")).Expect().Status(httptest.StatusOK)
	e.GET(("/departuresRegister/28/2019-01-01 00:00/2019-01-01 00:00/")).Expect().Status(httptest.StatusOK)
	e.GET(("/roadsteadRegister/28/2019-01-01 00:00/2019-01-01 00:00/")).Expect().Status(httptest.StatusOK)
	e.GET(("/mooredRegister/28/2019-01-01 00:00/2019-01-01 00:00/")).Expect().Status(httptest.StatusOK)

}
