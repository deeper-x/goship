# GSW - Goship [Shipreporting Webservice] 

Note: calls marked w/ [*MP] are marked for porting. Those w/ [*OK] are completed, deployed and available for production. Calls w/ [*SB] are in stand-by, candidated to be rejected.

## A - LIVE DATA SERVICES:

1. __Active trips__: [*SB]

```bash
http://<REMOTE_IP>:8000/activeTripsNow?id_portinformer=<id_portinformer>
```

2. __At roadstead__: [*OK]

```bash
http://<REMOTE_IP>:8000/anchored/<id_portinformer>
```

Result set:
```bash
# data := [idTrip, shipName, anchoringTime, currentActivity, anchoragePoint, shipType, iso3, grossTonnage, length, width, agency]

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
		}

```

3. __Moored__: [*OK]
    
```bash
http://<REMOTE_IP>:8000/moored/<id_portinformer>
```

Result set:
```bash
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
		}

```


4. __Arrivals__: [*OK]

```bash
http://<REMOTE_IP>:8000/arrivals/<id_portinformer>
```

Result set:
```bash
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


5. __Departures__: [*OK]

```bash
http://<REMOTE_IP>:8000/departures/<id_portinformer>
```
Result set:
```bash
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


6. __Arrival previsions__: [*OK]

```bash
http://<REMOTE_IP>:8000/arrivalPrevisionsToday/<id_portinformer>
```

Result set:
```bash
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

7. __Shipped goods__: [*OK]

```bash
http://<REMOTE_IP>:8000/shippedGoodToday/<id_portinformer>
```

Result set:
```bash
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

8. __RO/RO + RO/PAX__: [*OK]

```bash
http://<REMOTE_IP>:8000/trafficListToday/<id_portinformer>
```

Result set:
```bash
# data := [idTrip, shipName, numContainer, numPassengers, numCamion, numFurgoni, numRimorchi, numAuto, numMoto, numCamper, tons, numBus, numMinibus, trafficListMvntType, trafficListCategories, quay]


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

9. __Shifting previsions__: [*OK]

```bash
http://<REMOTE_IP>:8000/shiftingPrevisionsToday/<id_portinformer>
```

Result set:
```bash
# data := [ship, ts_shifting_prevision, ship_type, ship_flag, ship_width, ship_length, gross_tonnage,net_tonnage, draft_aft, draft_fwd, agency, starting_quay_berth, starting_roadstead, stop_quay_berth, stop_roadstead, cargo_on_board]

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
			"startingRoadstead":    startingRoadstead.String,
			"cargoOnBoard":         cargoOnBoard.String,
		}


```


10. __Departure previsions__: [*OK]
 
```bash
http://<REMOTE_IP>:8000/departurePrevisionsToday/<id_portinformer>
```

Result set:
```bash
# data := [ship, ts_departure_prevision, ship_type, ship_flag, ship_width, ship_length, gross_tonnage, net_tonnage, draft_aft, draft_fwd, agency, destination_port, starting_quay_berth, starting_roadstead, cargo_on_board]

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

## B - ARCHIVE DATA SERVICES:

1. __Trips archive [global recap, one row per trip]__: [*MP]

```bash
http://<REMOTE_IP>:8000/tripsArchive?id_portinformer=<ID_PORTINFORMER>
```

2. __Trips archive [global recap, one row per commercial operation]__: [*MP]

```bash
http://<REMOTE_IP>:8000/tripsArchiveMultiRows?id_portinformer=<ID_PORTINFORMER>
```

3. __Trip data archive__ [shipreport core]: [*SB]

```bash
http://<REMOTE_IP>:8000/shipReportList?id_portinformer=<ID_PORTINFORMER>
```

4. __Trip data archive detailed__ [shipreport]: [*SB]

```bash   
http://<REMOTE_IP>:8000/shipReportDetails?id_portinformer=<ID_PORTINFORMER>
```

5. __Arrivals archive__: [*MP]

```bash
http://<REMOTE_IP>:8000/arrivalsArchive?id_portinformer=<id_portinformer>
```

6. __Departures archive__: [*MP]

```bash
http://<REMOTE_IP>:8000/departuresArchive?id_portinformer=<id_portinformer>
```

7. __Shipped goods archive__: [*MP]

```bash
http://<REMOTE_IP>:8000/shippedGoodsArchive?id_portinformer=<id_portinformer>
```

8. __Traffic list archive__: [*MP]

```bash
http://<REMOTE_IP>:8000/trafficListArchive?id_portinformer=<id_portinformer>
```



## C - DAILY REGISTER SERVICES:

1. __Arrivals:__ [*MP]

```bash
http://<REMOTE_IP>:8000/registerArrivals?id_portinformer=<ID_PORTINFORMER>
```

2. __Moored:__ [*MP]

```bash
http://<REMOTE_IP>:8000/registerMoored?id_portinformer=<ID_PORTINFORMER>
```

3. __Roadstead:__ [*MP]
```bash
http://<REMOTE_IP>:8000/registerRoadstead?id_portinformer=<ID_PORTINFORMER>
```

4. __Departures:__ [*MP]
```bash
http://<REMOTE_IP>:8000/registerDepartures?id_portinformer=<ID_PORTINFORMER>
```

5. __Shiftings:__ [*MP]
```bash
http://<REMOTE_IP>:8000/registerShiftings?id_portinformer=<ID_PORTINFORMER>
```

6. __Arrival previsions:__ [*MP]
```bash
http://<REMOTE_IP>:8000/registerPlannedArrivals?id_portinformer=<ID_PORTINFORMER>
```

7. __Shipped goods:__ [*MP]
```bash
http://<REMOTE_IP>:8000/registerShippedGoods?id_portinformer=<ID_PORTINFORMER>
```

8. __RO/RO + RO/PAX:__ [*MP]
```bash
http://<REMOTE_IP>:8000/registerTrafficList?id_portinformer=<ID_PORTINFORMER>
```

## D - BUSINESS INTELLIGENCE SERVICES: ##

1. __Shiftings/maneuverings [per quay/berth]:__ [*MP]
```bash
http://<REMOTE_IP>:8000/tripsManeuverings?id_portinformer=<ID_PORTINFORMER>
```

2. __Shipped goods recap:__ [*MP]
```bash
http://<REMOTE_IP>:8000/shippedGoodsRecap?id_portinformer=<ID_PORTINFORMER>
```

3. __RO/RO + RO/PAX recap:__ [*MP]
```bash
http://<REMOTE_IP>:8000/trafficListRecap?id_portinformer=<ID_PORTINFORMER>
```

## E - METEO DATA ##
1. __Meteo data archive__:
```bash
http://<REMOTE_IP>:8000/meteoArchive?id_portinformer=<ID_PORTINFORMER>
```


# Install and (first) run 
```bash
# Install Go 1.11.9 and verify version
# ref. https://golang.org/doc/install?download=go1.11.9.linux-386.tar.gz
$ export PATH=${PATH}:/usr/local/go/bin/ # FIX w/ your installation path

# Pass vars at runtime or add to .bash_profile
$ export GOPATH=${HOME}/go
$ export GOBIN=${GOPATH}/bin
$ export PATH=${PATH}:${GOBIN}

$ go version
go version go1.11.9 linux/386

# Get goship and deploy
$ go get github.com/deeper-x/goship

$ cd ${GOPATH}/src/github.com/deeper-x/goship
$ go get -d ./...    
$ go install src/goship.go 
$ goship # RUN FOR TEST
Now listening on: http://localhost:8000
Application started. Press CTRL+C to shut down.

# Close test instance
$ <CTRL+C>  

```

# Systemd configuration

```bash
# Service install instructions: 
# This is a service file template you can start from 
$ sudo cat > /usr/lib/systemd/system/goship.service <<HEREDOC

[Unit]
Description=Shipreporting service middleware
Documentation=https://github.com/deeper-x/goship
After=network.target

[Service]
Type=simple
User=<YOUR_USER>
WorkingDirectory=<YOUR_GOPATH>/bin/
ExecStart=<YOUR_GOBIN>/goship
Restart=on-failure

[Install]
WantedBy=multi-user.target

HEREDOC
```


