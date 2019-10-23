package services

import (
	"github.com/deeper-x/goship/lib/db"
	"github.com/kataras/iris"
)

// ArrivalPrevisions todo description
func (objPortinformer Portinformer) ArrivalPrevisions(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	allArrivals := ldb.GetArrivalPrevisions(idPortinformer)
	ctx.JSON(allArrivals)
}

//DeparturePrevisions todo description
func (objPortinformer Portinformer) DeparturePrevisions(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	allDepartures := ldb.GetDeparturePrevisions(idPortinformer)
	ctx.JSON(allDepartures)
}

//ShiftingPrevisions todo description
func (objPortinformer Portinformer) ShiftingPrevisions(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	allShiftings := ldb.GetShiftingPrevisions(idPortinformer)
	ctx.JSON(allShiftings)
}
