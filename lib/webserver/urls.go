package webserver

import "github.com/deeper-x/goship/services"

var objPortinformer services.Portinformer

//URLLoader todo description
func (objInstance Instance) URLLoader() {
	objInstance.URLRequest("/moored/{id_portinformer:string}", objPortinformer.MooredNow)
}
