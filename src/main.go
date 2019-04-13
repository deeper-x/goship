package main

import (
	"github.com/deeper-x/goship/lib/webserver"
)

var app webserver.Instance

func main() {
	webserver.StartInstance(&app)

	app.PatternConfigurator()

	webserver.Run(&app)
}
