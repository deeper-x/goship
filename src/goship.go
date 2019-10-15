package main

import (
	"github.com/deeper-x/goship/lib/webserver"
)

// Inst object
var Inst webserver.Instance

func main() {
	//1. Start instance. Implicit to (&Inst).StartInstance()
	Inst.StartInstance()

	//2. Load url schema mapping
	Inst.URLLoader()

	//3. Run instance. Implicit to (&Inst).Run()
	Inst.Run()
}
