package services

import (
	"github.com/deeper-x/goship/lib/ldb"
	"github.com/kataras/iris"
)

// ArrivalPrevisions todo description
func (objPortinformer Portinformer) ArrivalPrevisions(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	conn := ldb.Connect()
	r := ldb.NewRepository(conn)

	allArrivals := r.GetArrivalPrevisions(idPortinformer)
	ctx.JSON(allArrivals)
}

//DeparturePrevisions todo description
func (objPortinformer Portinformer) DeparturePrevisions(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")

	conn := ldb.Connect()
	r := ldb.NewRepository(conn)

	allDepartures := r.GetDeparturePrevisions(idPortinformer)
	ctx.JSON(allDepartures)
}

//ShiftingPrevisions todo description
func (objPortinformer Portinformer) ShiftingPrevisions(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")

	conn := ldb.Connect()
	r := ldb.NewRepository(conn)

	allShiftings := r.GetShiftingPrevisions(idPortinformer)
	ctx.JSON(allShiftings)
}
