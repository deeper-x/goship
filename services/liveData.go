package services

import (
	"github.com/kataras/iris"
)

// DemoCall demo call
func DemoCall(ctx iris.Context) {
	ctx.HTML("hello, demo")
}
