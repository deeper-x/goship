package webserver

import "github.com/deeper-x/goship/services"

var objPortinformer services.Portinformer

//URLLoader todo description
func (objInstance Instance) URLLoader() {
	objInstance.URLRequest("/moored/{id_portinformer:string}", objPortinformer.MooredNow)
	objInstance.URLRequest("/anchored/{id_portinformer:string}", objPortinformer.RoadsteadNow)
	objInstance.URLRequest("/arrivalPrevisionsToday/{id_portinformer:string}", objPortinformer.ArrivalPrevisionsToday)
	objInstance.URLRequest("/departurePrevisionsToday/{id_portinformer:string}", objPortinformer.DeparturePrevisionsToday)
	objInstance.URLRequest("/shiftingPrevisionsToday/{id_portinformer:string}", objPortinformer.ShiftingPrevisionsToday)
	objInstance.URLRequest("/arrivals/{id_portinformer:string}", objPortinformer.Arrivals)
	objInstance.URLRequest("/departures/{id_portinformer:string}", objPortinformer.Departures)
	objInstance.URLRequest("/shippedGoodsToday/{id_portinformer:string}", objPortinformer.ShippedGoods)
}
