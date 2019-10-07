package webserver

import "github.com/deeper-x/goship/services"

var objPortinformer services.Portinformer

//URLLoader todo description
func (objInstance Instance) URLLoader() {
	// LIVE DATA CALLS:
	objInstance.URLRequest("/moored/{id_portinformer:string}", objPortinformer.MooredNow)      // verified
	objInstance.URLRequest("/anchored/{id_portinformer:string}", objPortinformer.RoadsteadNow) // verified
	objInstance.URLRequest("/arrivalPrevisionsToday/{id_portinformer:string}", objPortinformer.ArrivalPrevisionsToday)
	objInstance.URLRequest("/departurePrevisionsToday/{id_portinformer:string}", objPortinformer.DeparturePrevisionsToday)
	objInstance.URLRequest("/shiftingPrevisionsToday/{id_portinformer:string}", objPortinformer.ShiftingPrevisionsToday)
	objInstance.URLRequest("/arrivalsToday/{id_portinformer:string}", objPortinformer.ArrivalsToday)     // verified
	objInstance.URLRequest("/departuresToday/{id_portinformer:string}", objPortinformer.DeparturesToday) // verified
	objInstance.URLRequest("/shippedGoodsToday/{id_portinformer:string}", objPortinformer.ShippedGoods)
	objInstance.URLRequest("/trafficListToday/{id_portinformer:string}", objPortinformer.TrafficListToday)

	//ARCHIVE DATA CALLS
	objInstance.URLRequest("/arrivalsArchive/{id_portinformer:string}", objPortinformer.ArrivalsArchive)
	objInstance.URLRequest("/shippedGoodsArchive/{id_portinformer:string}", objPortinformer.ShippedGoodsArchive)
}
