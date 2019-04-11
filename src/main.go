package main

import (
	"github.com/deeper-x/goship/lib/webserver"
	"github.com/deeper-x/goship/services"
)

func main() {
	var app webserver.Instance

	webserver.StartInstance(&app)
	app.ManageRequest("/", services.DemoCall)

	webserver.Run(&app)

}
