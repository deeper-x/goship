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
