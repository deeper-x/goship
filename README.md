# GSW - Goship Shipreporting Webservice 

Note: calls marked w/ [*MP] are marked for porting. Those w/ [*OK] are completed, deployed and available for production. Calls w/ [*SB] are in stand-by, candidated to be rejected.

__Version: v0.2.12__

____

#### REST services:

Real time data:

- [A2 - Currently at roadstead](#a2-at-roadstead-ok): 

- [A3 - Currently at mooring](#a3-moored-ok)

- [A4 - Ships arrived today](#a4-arrivals-ok)

- [A5 - Ships departed today](#a5-departures-ok)

- [A6 - Today's arrivals previsions](#a6-arrival-previsions-ok)
  
- [A7 - Shipped goods for active trips](#a7-shipped-goods-ok)

- [A8 . Traffic list for active trips](#a8-roro--ropax-ok)

- [A9 - Today's shifting previsions](#a9-shifting-previsions-ok)

- [A10 - Today's departure previsions](#a10-departure-previsions-ok)

- [A11 - Today's shiftings](#a11-shiftings-ok)

Register data (with dynamic range):

- [C1 - Arrivals](#c1-arrivals-ok)
  
- [C2 - Moored](#c2-moored-ok)

- [C3 - At roadstead](#c3-roadstead-ok)

- [C4 - Departures](#c4-departures-ok)

- [C5 - Shiftings](#c5-shiftings-ok)

- [C7 - Shipped goods](#c7-shipped-goods-mp)

- [C8 - Traffic list](#c8-roro--ropax-mp)
  
Meteo data:

- [E2 - Active stations](#e2-active-meteo-stations)
  
  ____

#### A - LIVE DATA SERVICES:

#### A1. __Active trips__: [*SB] ##

Request:

```bash
http://<REMOTE_IP>:8000/activeTripsNow?id_portinformer=<id_portinformer>
```

#### A2. __At roadstead__: [*OK] ##

*Description:* trips currently in state of anchoring

Request:

```bash
Request:
http://<REMOTE_IP>:8000/anchored/<id_portinformer>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [idTrip, shipName, anchoringTime, currentActivity, anchoragePoint, shipType, iso3, grossTonnage, length, width, agency, shippedGoods, tsPlannedMooring, tsReadiness]

res := map[string]string{
    "id_trip":          idControlUnitDataStr.String,
    "ship":             shipName.String,
    "ship_type":        shipType.String,
    "anchoring_time":   anchoringTime.String,
    "current_activity": currentActivity.String,
    "anchorage_point":  anchoragePoint.String,
    "iso3":             iso3.String,
    "gross_tonnage":    grossTonnage.String,
    "length":           length.String,
    "width":            width.String,
    "agency":           agency.String,
    "shipped_goods":    shipped_goods.String,
    "ts_planned_mooring": tsPlannedMooring.String,
    "ts_readiness":       tsReadiness.String,
    }

```

#### A3. __Moored__: [*OK] ##

*Description:* trips currently in state of mooring

Request:

```bash
http://<REMOTE_IP>:8000/moored/<id_portinformer>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [idTrip, shipName, shipType, mooringTime, currentActivity, quay, shippedGoods, country, grossTonnage, length, width,  agency]

res := map[string]string{
    "id_trip":          idControlUnitDataStr.String,
    "ship":             shipName.String,
    "mooring_time":     mooringTime.String,
    "current_activity": currentActivity.String,
    "quay":             quay.String,
    "shipped_goods":    shippedGoods.String,
    "iso3":             iso3.String,
    "gross_tonnage":    grossTonnage.String,
    "ships_length":     length.String,
    "ships_width":      width.String,
    "ship_type":        shipType.String,
    "agency":           agency.String,
    "ts_etd":           tsETD.String,
}

```

#### A4. __Arrivals__: [*OK] ##

*Description:* trips with sighting date == ${TODAY}

Request:

```bash
http://<REMOTE_IP>:8000/arrivalsToday/<id_portinformer>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [ship, tsArrivalPrevision, shipType, country, width, length, grossTonnage, netTonnage, draftAft, draftFwd, agency, LPC, destinationQuayBerth, destinationRoadstead, cargoOnBoard]

res := map[string]string{
    "id_trip":                idTrip.String,
    "ship_name":              shipName.String,
    "ship_type":              shipType.String,
    "ts_sighting":            tsSighting.String,
    "ship_flag":              shipFlag.String,
    "ship_width":             shipWidth.String,
    "ship_length":            shipLength.String,
    "gross_tonnage":          grossTonnage.String,
    "net_tonnage":            netTonnage.String,
    "draft_aft":              draftAft.String,
    "draft_fwd":              draftFwd.String,
    "agency":                 agency.String,
    "last_port_of_call":      lastPortOfCall.String,
    "port_destination":       portDestination.String,
    "destination_quay_berth": destinationQuayBerth.String,
    "destination_roadstead":  destinationRoadstead.String,
    }


```

#### A5. __Departures__: [*OK] ##

*Description:* trips with out of sight date == ${TODAY}

Request:

```bash
http://<REMOTE_IP>:8000/departuresToday/<id_portinformer>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [idTrip, shipName, shipType, tsOutOfSight, shipFlag, shipWidth, shipLength, grossTonnage, netTonnage, draftAft, draftFwd, agency, LPC, portDestination]


res := map[string]string{
    "id_trip":           idTrip.String,
    "ship_name":         shipName.String,
    "ship_type":         shipType.String,
    "ts_out_of_sight":   tsOutOfSight.String,
    "ship_flag":         shipFlag.String,
    "ship_width":        shipWidth.String,
    "ship_length":       shipLength.String,
    "gross_tonnage":     grossTonnage.String,
    "net_tonnage":       netTonnage.String,
    "draft_aft":         draftAft.String,
    "draft_fwd":         draftFwd.String,
    "agency":            agency.String,
    "last_port_of_call": lastPortOfCall.String,
    "port_destination":  portDestination.String,
    }
```

#### A6. __Arrival previsions__: [*OK] ##

*Description:* trips with out_of_sight date == ${TODAY}

Request:

```bash
http://<REMOTE_IP>:8000/arrivalPrevisions/<id_portinformer>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [ship, tsArrivalPrevision, shipType, shipFlag, shipWidth, shipLength, grossTonnage, netTonnage, draftAft, draftFwd, agency, LPC, destinationQuayBerth, destinationRoadstead, cargoOnBoard]

res := map[string]string{
    "id_trip":                idControlUnitDataStr.String,
    "ship":                   shipName.String,
    "ts_arrival_prevision":   tsArrivalPrevision.String,
    "ship_type":              shipType.String,
    "ship_flag":              shipFlag.String,
    "ship_width":             shipWidth.String,
    "ship_length":            shipLength.String,
    "gross_tonnage":          grossTonnage.String,
    "net_tonnage":            netTonnage.String,
    "draft_aft":              draftAft.String,
    "draft_fwd":              draftFwd.String,
    "agency":                 agency.String,
    "last_port_of_call":      lastPortOfCall.String,
    "destination_quay_berth": destinationQuayBerth.String,
    "destination_roadstead":  destinationRoadstead.String,
    "cargo_on_board":         cargoOnBoard.String,
    }

```

#### A7. __Shipped goods__: [*OK] ##

*Description:* shipped goods in trips with sighting date == ${TODAY}

Request:

```bash
http://<REMOTE_IP>:8000/shippedGoodsToday/<id_portinformer>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [idTrip, shipName, quantity, unit, goodsCategory, shipType, shipFlag, shipWidth, shipLength, grossTonnage, netTonnage, groupCategory, macroCategory]

res := map[string]string{
    "id_trip":        idTrip.String,
    "ship_name":      shipName.String,
    "quantity":       quantity.String,
    "unit":           unit.String,
    "goods_category": goodsCategory.String,
    "ship_type":      shipType.String,
    "ship_flag":      shipFlag.String,
    "ship_width":     shipWidth.String,
    "shipLength":     shipLength.String,
    "grossTonnage":   grossTonnage.String,
    "netTonnage":     netTonnage.String,
    "groupCategory":  groupCategory.String,
    "macroCategory":  macroCategory.String,
    }

```

#### A8. __RO/RO + RO/PAX__: [*OK]

*Description:* RO/RO RO/PAX operations in trips with with sighting date == ${TODAY}

Request:

```bash
http://<REMOTE_IP>:8000/trafficListToday/<id_portinformer>
```

Response:

```bash
## Content-Type: application/json; charset=UTF-8

## data := [idTrip, shipName, numContainer, numPassengers, numCamion, numFurgoni, numRimorchi, numAuto, numMoto, numCamper, tons, numBus, numMinibus, trafficListMvntType, trafficListCategories, quay]


res := map[string]string{
    "id_trip":        idTrip.String,
    "ship_name":      shipName.String,
    "num_container":  numContainer.String,
    "num_passengers": numPassengers.String,
    "num_camion":     numCamion.String,
    "num_furgoni":    numFurgoni.String,
    "num_rimorchi":   numRimorchi.String,
    "num_auto":       numAuto.String,
    "num_moto":       numMoto.String,
    "num_camper":     numCamper.String,
    "tons":           tons.String,
    "num_bus":        numBus.String,
    "num_minibus":    numMinibus.String,
    "mvnt_type":      mvntType.String,
    "description":    description.String,
    "quay":           quay.String,
    }


```

#### A9. __Shifting previsions__: [*OK]

*Description:* trips with shifting prevision date == ${TODAY}

Request:

```bash
http://<REMOTE_IP>:8000/shiftingPrevisions/<id_portinformer>
```

Response:

```bash
## Content-Type: application/json; charset=UTF-8

## data := [ship, ts_shifting_prevision, ship_type, ship_flag, ship_width, ship_length, gross_tonnage,net_tonnage, draft_aft, draft_fwd, agency, starting_quay_berth, starting_roadstead, stop_quay_berth, stop_roadstead, cargo_on_board]


res := map[string]string{
    "ship":                 ship.String,
    "tsDeparturePrevision": tsShiftingPrevision.String,
    "shipType":             shipType.String,
    "shipFlag":             shipFlag.String,
    "shipWidth":            shipWidth.String,
    "shipLength":           shipLength.String,
    "grossTonnage":         grossTonnage.String,
    "netTonnage":           netTonnage.String,
    "draftAft":             draftAft.String,
    "draftFwd":             draftFwd.String,
    "agency":               agency.String,
    "destinationPort":      destinationPort.String,
    "startingQuayBerth":    startingQuayBerth.String,
    "startingRoadstead":    stopRoadstead.String,
    "stopQuayBerth":        stopQuayBerth.String,
    "stopRoadstead":        stopRoadstead.String,
    "cargoOnBoard":         cargoOnBoard.String,
    }


```

#### A10. __Departure previsions__: [*OK]

*Description:* trips with departure prevision's date == ${TODAY}

Request:

```bash
http://<REMOTE_IP>:8000/departurePrevisions/<id_portinformer>
```

Response:

```bash
## Content-Type: application/json; charset=UTF-8

## data := [ship, ts_departure_prevision, ship_type, ship_flag, ship_width, ship_length, gross_tonnage, net_tonnage, draft_aft, draft_fwd, agency, destination_port, starting_quay_berth, starting_roadstead, cargo_on_board]

res := map[string]string{
    "ship":                   ship.String,
    "ts_departure_prevision": tsDeparturePrevision.String,
    "ship_type":              shipType.String,
    "ship_flag":              shipFlag.String,
    "ship_width":             shipWidth.String,
    "ship_length":            shipLength.String,
    "gross_tonnage":          grossTonnage.String,
    "net_tonnage":            netTonnage.String,
    "draft_aft":              draftAft.String,
    "draft_fwd":              draftFwd.String,
    "agency":                 agency.String,
    "destination_port":       destinationPort.String,
    "starting_quay_berth":    startingQuayBerth.String,
    "starting_roadstead":     startingRoadstead.String,
    "cargo_on_board":         cargoOnBoard.String,
    }

```

#### A11. __Shiftings:__ [OK]

*Description:* trips' where starting shifting's date == ${TODAY}  

Request:

```bash
http://<REMOTE_IP>:8000/shiftingsToday/<ID_PORTINFORMER>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [idTrip, tsSighting, imo, ship, shipType, iso3, fromQuay, toQuay, fromAnch, toAnch]

res := map[string]string{
	"id_trip":     idTrip.String,
	"ts_sighting": tsSighting.String,
	"imo":         imo.String,
	"ship":        ship.String,
	"ship_type":   shipType.String,
	"iso3":        iso3.String,
	"from_quay":   fromQuay.String,
	"to_quay":     toQuay.String,
	"from_anch":   fromAnch.String,
	"to_anch":     toAnch.String,
}

```


____

#### B - ARCHIVE DATA SERVICES:

#### B1. __Trips archive [global recap, one row per trip]__: [*MP]

Request:

```bash
http://<REMOTE_IP>:8000/tripsArchive?id_portinformer=<ID_PORTINFORMER>
```

#### B2. __Trips archive [global recap, one row per commercial operation]__: [*SB]

Request:

```bash
http://<REMOTE_IP>:8000/tripsArchiveMultiRows?id_portinformer=<ID_PORTINFORMER>
```

#### B3. __Trip data archive__ [shipreport core]: [*SB]

Request:

```bash
http://<REMOTE_IP>:8000/shipReportList?id_portinformer=<ID_PORTINFORMER>
```

#### B4. __Trip data archive detailed__ [shipreport]: [*SB]

Request:

```bash
http://<REMOTE_IP>:8000/shipReportDetails?id_portinformer=<ID_PORTINFORMER>
```

#### B5. __Arrivals archive__: [*MP]

Request:

```bash
http://<REMOTE_IP>:8000/arrivalsArchive?id_portinformer=<id_portinformer>
```

#### B6. __Departures archive__: [*MP]

Request:

```bash
http://<REMOTE_IP>:8000/departuresArchive?id_portinformer=<id_portinformer>
```

#### B7. __Shipped goods archive__: [*MP]

Request:

```bash
http://<REMOTE_IP>:8000/shippedGoodsArchive?id_portinformer=<id_portinformer>
```

#### B8. __Traffic list archive__: [*MP]

Request:

```bash
http://<REMOTE_IP>:8000/trafficListArchive?id_portinformer=<id_portinformer>
```

____

#### C - DAILY REGISTER SERVICES:

These calls get the last trip activity in range, according with criteria

#### C1. __Arrivals:__ [*OK]

*Description:* trips with sighting's date in range

Request:

```bash
http://<REMOTE_IP>:8000/arrivalsRegister/<ID_PORTINFORMER>/<TIMESTAMP_START>/<TIMESTAMP_STOP>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [ship, tsArrivalPrevision, shipType, country, width, length, grossTonnage, netTonnage, draftAft, draftFwd, agency, LPC, destinationQuayBerth, destinationRoadstead, cargoOnBoard]

res := map[string]string{
    "id_trip":                idTrip.String,
    "ship_name":              shipName.String,
    "ship_type":              shipType.String,
    "ts_sighting":            tsSighting.String,
    "ship_flag":              shipFlag.String,
    "ship_width":             shipWidth.String,
    "ship_length":            shipLength.String,
    "gross_tonnage":          grossTonnage.String,
    "net_tonnage":            netTonnage.String,
    "draft_aft":              draftAft.String,
    "draft_fwd":              draftFwd.String,
    "agency":                 agency.String,
    "last_port_of_call":      lastPortOfCall.String,
    "port_destination":       portDestination.String,
    "destination_quay_berth": destinationQuayBerth.String,
    "destination_roadstead":  destinationRoadstead.String,
    }


```

#### C2. __Moored:__ [*OK]

*Description:* trips with the last activity as mooring in range (whatever form has been used)

Request:

```bash
http://<REMOTE_IP>:8000/mooredRegister/<ID_PORTINFORMER>/<TIMESTAMP_START>/<TIMESTAMP_STOP>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [idTrip, shipName, shipType, tsMooring, shipFlag, shipWidth, shipLength, grossTonnage, netTonnage, agency]

res := map[string]string{
    "id_trip":       idTrip.String,
    "ship_name":     shipName.String,
    "ship_type":     shipType.String,
    "ts_mooring":    tsMooring.String,
    "ship_flag":     shipFlag.String,
    "ship_width":    shipWidth.String,
    "ship_length":   shipLength.String,
    "gross_tonnage": grossTonnage.String,
    "net_tonnage":   netTonnage.String,
    "agency":        agency.String,
    }
```

#### C3. __Roadstead:__ [*OK]

*Description:* trips with the last activity as anchoring in range (whatever form has been used)

Request:

```bash
http://<REMOTE_IP>:8000/roadsteadRegister/<ID_PORTINFORMER>/<TIMESTAMP_START>/<TIMESTAMP_STOP>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [idTrip, shipName, shipType, tsAnchoring, shipFlag, shipWidth, shipLength, grossTonnage, netTonnage, agency]

res := map[string]string{
    "id_trip":       idTrip.String,
    "ship_name":     shipName.String,
    "ship_type":     shipType.String,
    "ts_anchoring":  tsAnchoring.String,
    "ship_flag":     shipFlag.String,
    "ship_width":    shipWidth.String,
    "ship_length":   shipLength.String,
    "gross_tonnage": grossTonnage.String,
    "net_tonnage":   netTonnage.String,
    "agency":        agency.String,
    }
```

#### C4. __Departures:__ [*OK]

*Description:* trips with out of sight in range (whatever form has been used)

Request:

```bash
http://<REMOTE_IP>:8000/departuresRegister/<ID_PORTINFORMER>/<TIMESTAMP_START>/<TIMESTAMP_STOP>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [idTrip, shipName, shipType, tsOutOfSight, shipFlag, shipWidth, shipLength, grossTonnage, netTonnage, draftAft, draftFwd, agency, LPC, portDestination]


res := map[string]string{
    "id_trip":           idTrip.String,
    "ship_name":         shipName.String,
    "ship_type":         shipType.String,
    "ts_out_of_sight":   tsOutOfSight.String,
    "ship_flag":         shipFlag.String,
    "ship_width":        shipWidth.String,
    "ship_length":       shipLength.String,
    "gross_tonnage":     grossTonnage.String,
    "net_tonnage":       netTonnage.String,
    "draft_aft":         draftAft.String,
    "draft_fwd":         draftFwd.String,
    "agency":            agency.String,
    "last_port_of_call": lastPortOfCall.String,
    "port_destination":  portDestination.String,
    }
```

#### C5. __Shiftings:__ [OK]

*Description:* trips' starting shifting in passed range  

Request:

```bash
http://<REMOTE_IP>:8000/shiftingsRegister/<ID_PORTINFORMER>/<TIMESTAMP_START>/<TIMESTAMP_STOP>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [idTrip, tsSighting, imo, ship, shipType, iso3, fromQuay, toQuay, fromAnch, toAnch]

res := map[string]string{
	"id_trip":     idTrip.String,
	"ts_sighting": tsSighting.String,
	"imo":         imo.String,
	"ship":        ship.String,
	"ship_type":   shipType.String,
	"iso3":        iso3.String,
	"from_quay":   fromQuay.String,
	"to_quay":     toQuay.String,
	"from_anch":   fromAnch.String,
	"to_anch":     toAnch.String,
}

```

#### C6. __Arrival previsions:__ [*MP]

Request:

```bash
http://<REMOTE_IP>:8000/registerPlannedArrivals?id_portinformer=<ID_PORTINFORMER>
```

#### C7. __Shipped goods:__ [*MP]

*Description:* commercial operations trips with sighting date in range (whatever form has been used)

Request:

```bash
http://<REMOTE_IP>:8000/shippedGoodsRegister/<id_portinformer>/<TIMESTAMP_START>/<TIMESTAMP_STOP>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [idTrip, shipName, quantity, unit, goodsCategory, shipType, shipFlag, shipWidth, shipLength, grossTonnage, netTonnage, groupCategory, macroCategory]

res := map[string]string{
    "id_trip":        idTrip.String,
    "ship_name":      shipName.String,
    "quantity":       quantity.String,
    "unit":           unit.String,
    "goods_category": goodsCategory.String,
    "ship_type":      shipType.String,
    "ship_flag":      shipFlag.String,
    "ship_width":     shipWidth.String,
    "ship_length":    shipLength.String,
    "gross_tonnage":  grossTonnage.String,
    "net_tonnage":    netTonnage.String,
    "group_category": groupCategory.String,
    "macro_category": macroCategory.String,
}
```

#### C8. __RO/RO + RO/PAX:__ [*MP]

*Description:* RO/RO + RO/PAX operations trips with sighting date in range (whatever form has been used)


Request:

```bash
http://<REMOTE_IP>:8000/trafficListRegister/<id_portinformer>/<TIMESTAMP_START>/<TIMESTAMP_STOP>
```

Response:

```bash
## Content-Type: application/json; charset=UTF-8

## data := [idTrip, shipName, tsSighting, numContainer, numPassengers, numCamion, numFurgoni, numRimorchi, numAuto, numMoto, numCamper, tons, numBus, numMinibus, trafficListMvntType, trafficListCategories, quay]


res := map[string]string{
    "id_trip":        idTrip.String,
    "ship_name":      shipName.String,
    "ts_sighting":    tsSighting.String,
    "num_container":  numContainer.String,
    "num_passengers": numPassengers.String,
    "num_camion":     numCamion.String,
    "num_furgoni":    numFurgoni.String,
    "num_rimorchi":   numRimorchi.String,
    "num_auto":       numAuto.String,
    "num_moto":       numMoto.String,
    "num_camper":     numCamper.String,
    "tons":           tons.String,
    "num_bus":        numBus.String,
    "num_minibus":    numMinibus.String,
    "mvnt_type":      mvntType.String,
    "description":    description.String,
    "quay":           quay.String,
    }


```
____

#### D - BUSINESS INTELLIGENCE SERVICES: ####

#### D1. __Shiftings/maneuverings [per quay/berth]:__ [*MP] ####

Request:

```bash
http://<REMOTE_IP>:8000/tripsManeuverings?id_portinformer=<ID_PORTINFORMER>
```

#### D2. __Shipped goods recap:__ [*MP] ####

Request:

```bash
http://<REMOTE_IP>:8000/shippedGoodsRecap?id_portinformer=<ID_PORTINFORMER>
```

#### D3. __RO/RO + RO/PAX recap:__ [*MP]

Request:

```bash
http://<REMOTE_IP>:8000/trafficListRecap?id_portinformer=<ID_PORTINFORMER>
```
____

#### E - METEO DATA ####


#### E1. __Meteo data archive__: [*MP]

Request:

```bash
http://<REMOTE_IP>:8000/meteoArchive?id_portinformer=<ID_PORTINFORMER>
```

#### E2. __Active meteo stations__: [OK]

Request:

```bash
http://<REMOTE_IP>:8000/weatherActiveStations
```

Response:

```bash
## Content-Type: application/json; charset=UTF-8

## data := [idPortinformer, portinformerCode, tsFirstCreated, isActive]

res := map[string]string{
    "id_portinformer":   idPortinformer.String,
    "portinformer_code": portinformerCode.String,
    "ts_first_created":  tsFirstCreated.String,
    "is_active":         isActive.String,
}
```

____


#### Install Go and (first) run 
ref. https://golang.org/doc/install?download=go1.12.10.linux-386.tar.gz

```bash
$ go version
go version go1.12.10 linux/amd64

$ export PATH=${PATH}:/usr/local/go/bin/ ## FIX w/ your installation path

## Pass vars at runtime or add to .bash_profile
$ export GOPATH=${HOME}/go
$ export GOBIN=${GOPATH}/bin
$ export PATH=${PATH}:${GOBIN}

## Get goship and deploy
$ go get github.com/deeper-x/goship

$ cd ${GOPATH}/src/github.com/deeper-x/goship
$ cat > .env <<HEREDOC 
DB_DSN="postgres://<user>:<passwd>@127.0.0.1/<db>"
HEREDOC

$ go get -d ./...    
$ go install src/goship.go 
$ goship ## RUN FOR TEST
Now listening on: http://localhost:8000
Application started. Press CTRL+C to shut down.

## Close test instance
$ <CTRL+C>  

```

____

#### Integration Test + Unit test
```bash
$ go test -v ./...
?   	github.com/deeper-x/goship/conf	[no test files]
=== RUN   TestAllRoadstead
--- PASS: TestAllRoadstead (0.01s)
=== RUN   TestAllMoored
--- PASS: TestAllMoored (0.01s)
=== RUN   TestArrivalsToday
--- PASS: TestArrivalsToday (0.01s)
=== RUN   TestDeparturesToday
--- PASS: TestDeparturesToday (0.01s)
=== RUN   TestArrivalPrevisionsToday
--- PASS: TestArrivalPrevisionsToday (0.01s)
=== RUN   TestShippedGoods
--- PASS: TestShippedGoods (0.01s)
=== RUN   TestTrafficList
--- PASS: TestTrafficList (0.01s)
=== RUN   TestShiftingPrevisionsToday
--- PASS: TestShiftingPrevisionsToday (0.00s)
=== RUN   TestDeparturePrevisionsToday
--- PASS: TestDeparturePrevisionsToday (0.01s)
=== RUN   TestShiftingsToday
--- PASS: TestShiftingsToday (0.01s)
=== RUN   TestArrivalsRegister
--- PASS: TestArrivalsRegister (0.00s)
=== RUN   TestMooredRegister
--- PASS: TestMooredRegister (0.00s)
=== RUN   TestAnchoredRegister
--- PASS: TestAnchoredRegister (0.01s)
=== RUN   TestDeparturesRegister
--- PASS: TestDeparturesRegister (0.00s)
=== RUN   TestShiftingsRegister
--- PASS: TestShiftingsRegister (0.01s)
=== RUN   TestShippedGoodsRegister
--- PASS: TestShippedGoodsRegister (0.00s)
=== RUN   TestTrafficListRegister
--- PASS: TestTrafficListRegister (0.00s)
PASS
ok  	github.com/deeper-x/goship/lib/ldb	0.132s
=== RUN   TestLive
--- PASS: TestLive (0.41s)
=== RUN   TestRegister
--- PASS: TestRegister (0.24s)
=== RUN   TestMeteo
--- PASS: TestMeteo (0.01s)
PASS
ok  	github.com/deeper-x/goship/src	0.710s
deeper-x@local-PC:~/go/src/github.com/deeper-x/goship$ 


```

____

#### Systemd configuration

```bash
## Service install instructions: 
## This is a service file template you can start from 
$ sudo cat > /usr/lib/systemd/system/goship.service <<HEREDOC

[Unit]
Description=Shipreporting service middleware
Documentation=https://github.com/deeper-x/goship
After=network.target

[Service]
Type=simple
User=<YOUR_USER>
Environment=GOPATH=/home/<YOUR_USER>/go
WorkingDirectory=<YOUR_GOPATH>/bin/
ExecStart=<YOUR_GOBIN>/goship
Restart=on-failure

[Install]
WantedBy=multi-user.target

HEREDOC
```


