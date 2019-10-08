package webserver

import "github.com/deeper-x/goship/services"

var objPortinformer services.Portinformer

//URLLoader todo description
func (objInstance Instance) URLLoader() {
	// LIVE DATA CALLS:
	objInstance.URLRequest("/moored/{id_portinformer:string}", objPortinformer.MooredNow)                        // verified
	objInstance.URLRequest("/anchored/{id_portinformer:string}", objPortinformer.RoadsteadNow)                   // verified
	objInstance.URLRequest("/arrivalPrevisions/{id_portinformer:string}", objPortinformer.ArrivalPrevisions)     // verified
	objInstance.URLRequest("/departurePrevisions/{id_portinformer:string}", objPortinformer.DeparturePrevisions) // verified
	objInstance.URLRequest("/shiftingPrevisions/{id_portinformer:string}", objPortinformer.ShiftingPrevisions)   // verified
	objInstance.URLRequest("/arrivalsToday/{id_portinformer:string}", objPortinformer.ArrivalsToday)             // verified
	objInstance.URLRequest("/departuresToday/{id_portinformer:string}", objPortinformer.DeparturesToday)         // verified
	objInstance.URLRequest("/shippedGoodsToday/{id_portinformer:string}", objPortinformer.ShippedGoods)
	objInstance.URLRequest("/trafficListToday/{id_portinformer:string}", objPortinformer.TrafficListToday)

	//REGISTER DATA CALLS
	objInstance.URLRequest("/arrivalsRegister/{id_portinformer:string}/{start:string}/{stop:string}", objPortinformer.ArrivalsRegister)
	objInstance.URLRequest("/departuresRegister/{id_portinformer:string}/{start:string}/{stop:string}", objPortinformer.DeparturesRegister)

	//ARCHIVE DATA CALLS
	objInstance.URLRequest("/arrivalsArchive/{id_portinformer:string}", objPortinformer.ArrivalsArchive)
	objInstance.URLRequest("/shippedGoodsArchive/{id_portinformer:string}", objPortinformer.ShippedGoodsArchive)
}
