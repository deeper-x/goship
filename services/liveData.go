package services

import (
	"github.com/deeper-x/goship/lib/ldb"
	"github.com/kataras/iris"
)

// MooredNow demo call
func (objPortinformer Portinformer) MooredNow(ctx iris.Context) {
	idPortinformer := ctx.Params().Get("id_portinformer")
	allMoored := ldb.GetAllMoored(idPortinformer)
	ctx.JSON(allMoored)
}
