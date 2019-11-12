package services

import (
	"github.com/deeper-x/goship/lib/ldb"
	"github.com/kataras/iris/v12"
)

// Home todo doc
func (objPortinformer Portinformer) Home(ctx iris.Context) {
	ctx.JSON("Please choose a service")
}

// MooredNow todo description
func (objPortinformer Portinformer) MooredNow(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")

	conn := ldb.Connect()
	r := ldb.NewRepository(conn)

	allMoored := r.GetAllMoored(idPortinformer)

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(allMoored)
}

// RoadsteadNow todo description
func (objPortinformer Portinformer) RoadsteadNow(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")

	conn := ldb.Connect()
	r := ldb.NewRepository(conn)

	allAnchoring := r.GetAllRoadstead(idPortinformer)

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(allAnchoring)
}

// ArrivalsToday todo description
func (objPortinformer Portinformer) ArrivalsToday(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")

	conn := ldb.Connect()
	r := ldb.NewRepository(conn)

	arrivals := r.GetTodayArrivals(idPortinformer, 10)

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(arrivals)
}

// DeparturesToday todo description
func (objPortinformer Portinformer) DeparturesToday(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")

	conn := ldb.Connect()
	r := ldb.NewRepository(conn)

	departures := r.GetTodayDepartures(idPortinformer, 26)

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(departures)
}

// ShippedGoods todo description
func (objPortinformer Portinformer) ShippedGoods(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")

	conn := ldb.Connect()
	r := ldb.NewRepository(conn)

	shippedGoods := r.GetTodayShippedGoods(idPortinformer)

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(shippedGoods)
}

// ShiftingsToday todo description
func (objPortinformer Portinformer) ShiftingsToday(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	conn := ldb.Connect()
	r := ldb.NewRepository(conn)

	shiftings := r.GetTodayShiftings(idPortinformer)

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(shiftings)
}

//TrafficListToday todo description
func (objPortinformer Portinformer) TrafficListToday(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	conn := ldb.Connect()
	r := ldb.NewRepository(conn)

	trafficList := r.GetTodayTrafficList(idPortinformer)

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(trafficList)
}
