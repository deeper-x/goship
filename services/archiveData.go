package services

import (
	"github.com/deeper-x/goship/lib/ldb"
	"github.com/kataras/iris"
)

//ArrivalsArchive todo desc
func (objPortinformer Portinformer) ArrivalsArchive(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	allArrivals := ldb.GetAllArrivalsArchive(idPortinformer, 10)
	ctx.JSON(allArrivals)
}
