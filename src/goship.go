package main

import (
	"github.com/deeper-x/goship/lib/webserver"
)

// Inst object
var Inst webserver.Instance

func main() {
	webserver.StartInstance(&Inst)
	Inst.URLLoader()
	webserver.Run(&Inst)
}
