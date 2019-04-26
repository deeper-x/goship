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
