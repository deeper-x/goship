package services

import (
	"github.com/deeper-x/goship/lib/db"
	"github.com/kataras/iris"
)

// ActiveStations todo doc
func (objPortinformer Portinformer) ActiveStations(ctx iris.Context) {
	activeStations := ldb.GetActiveStations()
	ctx.JSON(activeStations)
}
