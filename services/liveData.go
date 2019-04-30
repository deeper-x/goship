package services

import (
	"github.com/deeper-x/goship/lib/ldb"
	"github.com/kataras/iris"
)

// MooredNow todo description
func (objPortinformer Portinformer) MooredNow(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	allMoored := ldb.GetAllMoored(idPortinformer)
	ctx.JSON(allMoored)
}

// RoadsteadNow todo description
func (objPortinformer Portinformer) RoadsteadNow(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	allAnchoring := ldb.GetAllRoadstead(idPortinformer)
	ctx.JSON(allAnchoring)
}

// ArrivalPrevisions todo description
func (objPortinformer Portinformer) ArrivalPrevisions(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	allArrivals := ldb.GetActiveArrivalPrevisions(idPortinformer)
	ctx.JSON(allArrivals)
}

// Arrivals todo description
func (objPortinformer Portinformer) Arrivals(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	arrivals := ldb.GetTodayArrivals(idPortinformer, "10")
	ctx.JSON(arrivals)
}

// Departures todo description
func (objPortinformer Portinformer) Departures(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	departures := ldb.GetTodayDepartures(idPortinformer, 26)
	ctx.JSON(departures)
}
