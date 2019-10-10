package services

import (
	"github.com/deeper-x/goship/lib/ldb"
	"github.com/kataras/iris"
)

// Home todo doc
func (objPortinformer Portinformer) Home(ctx iris.Context) {
	ctx.JSON("hello, world")
}

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

// ArrivalsToday todo description
func (objPortinformer Portinformer) ArrivalsToday(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	arrivals := ldb.GetTodayArrivals(idPortinformer, 10)
	ctx.JSON(arrivals)
}

// DeparturesToday todo description
func (objPortinformer Portinformer) DeparturesToday(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	departures := ldb.GetTodayDepartures(idPortinformer, 26)
	ctx.JSON(departures)
}

// ShippedGoods todo description
func (objPortinformer Portinformer) ShippedGoods(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	shippedGoods := ldb.GetTodayShippedGoods(idPortinformer)
	ctx.JSON(shippedGoods)
}

//TrafficListToday todo description
func (objPortinformer Portinformer) TrafficListToday(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	trafficList := ldb.GetTodayTrafficList(idPortinformer)
	ctx.JSON(trafficList)
}
