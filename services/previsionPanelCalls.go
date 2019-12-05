package services

import (
	"github.com/deeper-x/goship/lib/ldb"
	"github.com/kataras/iris/v12"
)

// ArrivalPrevisions todo description
func (objPortinformer Portinformer) ArrivalPrevisions(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	conn := ldb.Connect()
	r := ldb.NewRepository(conn)

	allArrivals := r.GetArrivalPrevisions(idPortinformer)
	conn.Close()

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(allArrivals)
}

//DeparturePrevisions todo description
func (objPortinformer Portinformer) DeparturePrevisions(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")

	conn := ldb.Connect()
	r := ldb.NewRepository(conn)

	allDepartures := r.GetDeparturePrevisions(idPortinformer)
	conn.Close()

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(allDepartures)
}

//ShiftingPrevisions todo description
func (objPortinformer Portinformer) ShiftingPrevisions(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")

	conn := ldb.Connect()
	r := ldb.NewRepository(conn)

	allShiftings := r.GetShiftingPrevisions(idPortinformer)
	conn.Close()

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(allShiftings)
}
