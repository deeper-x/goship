package services

import (
	"github.com/deeper-x/goship/lib/ldb"
	"github.com/kataras/iris"
)

// ArrivalPrevisionsToday todo description
func (objPortinformer Portinformer) ArrivalPrevisionsToday(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	allArrivals := ldb.GetTodayArrivalPrevisions(idPortinformer)
	ctx.JSON(allArrivals)
}

//DeparturePrevisionsToday todo description
func (objPortinformer Portinformer) DeparturePrevisionsToday(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	allDepartures := ldb.GetTodayDeparturePrevisions(idPortinformer)
	ctx.JSON(allDepartures)
}

//ShiftingPrevisionsToday todo description
func (objPortinformer Portinformer) ShiftingPrevisionsToday(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	allShiftings := ldb.GetTodayShiftingPrevisions(idPortinformer)
	ctx.JSON(allShiftings)
}
