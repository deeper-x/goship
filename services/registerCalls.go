package services

import (
	"github.com/deeper-x/goship/lib/ldb"
	"github.com/kataras/iris"
)

// ArrivalsRegister todo doc
func (objPortinformer Portinformer) ArrivalsRegister(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	start := ctx.Params().Get("start")
	stop := ctx.Params().Get("stop")

	allArrivals := ldb.GetArrivalsRegister(idPortinformer, 10, start, stop)
	ctx.JSON(allArrivals)
}

// DeparturesRegister todo doc
func (objPortinformer Portinformer) DeparturesRegister(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	start := ctx.Params().Get("start")
	stop := ctx.Params().Get("stop")

	allDepartures := ldb.GetDeparturesRegister(idPortinformer, 26, start, stop)
	ctx.JSON(allDepartures)
}

// RoadsteadRegister todo doc
func (objPortinformer Portinformer) RoadsteadRegister(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	start := ctx.Params().Get("start")
	stop := ctx.Params().Get("stop")

	allRoadstead := ldb.GetRoadsteadRegister(idPortinformer, start, stop)
	ctx.JSON(allRoadstead)
}

// MooredRegister todo doc
func (objPortinformer Portinformer) MooredRegister(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	start := ctx.Params().Get("start")
	stop := ctx.Params().Get("stop")

	allMoored := ldb.GetMooredRegister(idPortinformer, start, stop)
	ctx.JSON(allMoored)
}

// ShiftingsRegister todo doc
func (objPortinformer Portinformer) ShiftingsRegister(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	start := ctx.Params().Get("start")
	stop := ctx.Params().Get("stop")

	allShiftings := ldb.GetShiftingsRegister(idPortinformer, start, stop)
	ctx.JSON(allShiftings)
}

// ShippedGoodsRegister todo doc
func (objPortinformer Portinformer) ShippedGoodsRegister(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	start := ctx.Params().Get("start")
	stop := ctx.Params().Get("stop")

	allShippedGoods := ldb.GetShippedGoodsRegister(idPortinformer, start, stop)

	ctx.JSON(allShippedGoods)
}

//TrafficListRegister todo description
func (objPortinformer Portinformer) TrafficListRegister(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	start := ctx.Params().Get("start")
	stop := ctx.Params().Get("stop")

	trafficList := ldb.GetRegisterTrafficList(idPortinformer, start, stop)
	ctx.JSON(trafficList)
}
