package webserver

import "github.com/deeper-x/goship/services"

var objPortinformer services.Portinformer

//URLLoader todo description
func (objInstance Instance) URLLoader() {
	objInstance.URLRequest("/moored/{id_portinformer:string}", objPortinformer.MooredNow)
	objInstance.URLRequest("/anchored/{id_portinformer:string}", objPortinformer.RoadsteadNow)
	objInstance.URLRequest("/arrivals/{id_portinformer:string}", objPortinformer.ArrivalsNow)
}
