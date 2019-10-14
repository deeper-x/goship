package ldb

import (
	"database/sql"
	"log"
)

var odot = GenDot("./qsql/live_queries.sql")

// GetAllRoadstead todo doc
func GetAllRoadstead(idPortinformer string) []map[string]string {
	var idControlUnitData sql.NullString
	var shipName, anchoringTime, currentActivity, shippedGoods sql.NullString
	var shipType, iso3, grossTonnage, anchoragePoint, length, width, agency sql.NullString
	var result []map[string]string = []map[string]string{}

	connector := Connect()

	rows, err := odot.Query(connector, "all-anchored", idPortinformer, idPortinformer)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&idControlUnitData,
			&shipName,
			&anchoringTime,
			&currentActivity,
			&anchoragePoint,
			&shipType,
			&iso3,
			&grossTonnage,
			&length,
			&width,
			&agency,
			&shippedGoods,
		)

		if err != nil {
			log.Fatal(err)
		}

		idControlUnitDataStr := idControlUnitData

		tmpDict := map[string]string{
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
			"shipped_goods":    shippedGoods.String,
		}

		result = append(result, tmpDict)
	}

	return result
}

//GetArrivalPrevisions todo doc
func GetArrivalPrevisions(idPortinformer string) []map[string]string {
	var idControlUnitData, shipName sql.NullString
	var tsArrivalPrevision, shipType sql.NullString
	var shipFlag, shipWidth, shipLength, grossTonnage sql.NullString
	var netTonnage, draftAft, draftFwd sql.NullString
	var agency, cargoOnBoard sql.NullString
	var lastPortOfCall sql.NullString
	var destinationQuayBerth sql.NullString
	var destinationRoadstead sql.NullString

	var result []map[string]string = []map[string]string{}

	connector := Connect()

	rows, err := odot.Query(connector, "arrival-previsions", idPortinformer)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&shipName,
			&tsArrivalPrevision,
			&shipType,
			&shipFlag,
			&shipWidth,
			&shipLength,
			&grossTonnage,
			&netTonnage,
			&draftAft,
			&draftFwd,
			&agency,
			&lastPortOfCall,
			&destinationQuayBerth,
			&destinationRoadstead,
			&cargoOnBoard,
		)

		if err != nil {
			log.Fatal(err)
		}

		idControlUnitDataStr := idControlUnitData

		tmpDict := map[string]string{
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

		result = append(result, tmpDict)
	}

	return result
}

//GetShiftingPrevisions todo doc
func GetShiftingPrevisions(idPortinformer string) []map[string]string {
	var ship, tsShiftingPrevision, shipType, shipFlag, shipWidth sql.NullString
	var shipLength, grossTonnage, netTonnage, draftAft, draftFwd sql.NullString
	var agency, destinationPort, startingQuayBerth, startingRoadstead, stopQuayBerth, stopRoadstead, cargoOnBoard sql.NullString

	var result []map[string]string = []map[string]string{}

	connector := Connect()

	rows, err := odot.Query(connector, "shifting-previsions", idPortinformer)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(
			&ship,
			&tsShiftingPrevision,
			&shipType,
			&shipFlag,
			&shipWidth,
			&shipLength,
			&grossTonnage,
			&netTonnage,
			&draftAft,
			&draftFwd,
			&agency,
			&startingQuayBerth,
			&stopQuayBerth,
			&startingRoadstead,
			&stopRoadstead,
			&cargoOnBoard,
		)

		if err != nil {
			log.Fatal(err)
		}

		tmpDict := map[string]string{
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

		result = append(result, tmpDict)
	}

	return result
}

//GetDeparturePrevisions todo doc
func GetDeparturePrevisions(idPortinformer string) []map[string]string {
	var ship, tsDeparturePrevision, shipType, shipFlag, shipWidth sql.NullString
	var shipLength, grossTonnage, netTonnage, draftAft, draftFwd sql.NullString
	var agency, destinationPort, startingQuayBerth, startingRoadstead, cargoOnBoard sql.NullString

	var result []map[string]string = []map[string]string{}

	connector := Connect()

	rows, err := odot.Query(connector, "departure-previsions", idPortinformer)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&ship,
			&tsDeparturePrevision,
			&shipType,
			&shipFlag,
			&shipWidth,
			&shipLength,
			&grossTonnage,
			&netTonnage,
			&draftAft,
			&draftFwd,
			&agency,
			&destinationPort,
			&startingQuayBerth,
			&startingRoadstead,
			&cargoOnBoard,
		)

		if err != nil {
			log.Fatal(err)
		}

		tmpDict := map[string]string{
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

		result = append(result, tmpDict)
	}

	return result
}

//GetAllMoored todo doc
func GetAllMoored(idPortinformer string) []map[string]string {
	var idControlUnitData, iso3, grossTonnage, length, width, shipType sql.NullString
	var shipName, mooringTime, currentActivity, quay, shippedGoods sql.NullString
	var agency sql.NullString

	var result []map[string]string

	connector := Connect()

	rows, err := odot.Query(connector, "all-moored", idPortinformer, idPortinformer)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&idControlUnitData,
			&shipName,
			&mooringTime,
			&currentActivity,
			&quay,
			&shippedGoods,
			&iso3,
			&grossTonnage,
			&length,
			&width,
			&shipType,
			&agency,
		)

		if err != nil {
			log.Fatal(err)
		}

		idControlUnitDataStr := idControlUnitData

		tmpDict := map[string]string{
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
		result = append(result, tmpDict)
	}

	return result
}

//GetTodayArrivals todo doc
func GetTodayArrivals(idPortinformer string, idArrivalPrevision int) []map[string]string {
	var idTrip, shipName, shipType, tsSighting, shipFlag, shipWidth, shipLength sql.NullString
	var grossTonnage, netTonnage, draftAft, draftFwd, agency, lastPortOfCall sql.NullString
	var portDestination, destinationQuayBerth, destinationRoadstead sql.NullString

	var result []map[string]string = []map[string]string{}

	connector := Connect()

	rows, err := odot.Query(connector, "arrivals", idArrivalPrevision, idArrivalPrevision, idArrivalPrevision, idArrivalPrevision, idPortinformer)

	if err != nil {
		log.Fatal(err)
	}

	defer connector.Close()

	for rows.Next() {
		err := rows.Scan(
			&idTrip,
			&shipName,
			&shipType,
			&tsSighting,
			&shipFlag,
			&shipWidth,
			&shipLength,
			&grossTonnage,
			&netTonnage,
			&draftAft,
			&draftFwd,
			&agency,
			&lastPortOfCall,
			&portDestination,
			&destinationQuayBerth,
			&destinationRoadstead,
		)

		if err != nil {
			log.Fatal(err)
		}

		tmpDict := map[string]string{
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

		result = append(result, tmpDict)
	}

	return result
}

// GetTodayDepartures todo description
func GetTodayDepartures(idPortinformer string, idDepartureState int) []map[string]string {
	var idTrip, shipName, shipType, tsOutOfSight, shipFlag, shipWidth sql.NullString
	var shipLength, grossTonnage sql.NullString
	var netTonnage, draftAft, draftFwd, agency, lastPortOfCall, portDestination sql.NullString

	var result []map[string]string = []map[string]string{}

	connector := Connect()

	rows, err := odot.Query(connector, "departures", idDepartureState, idPortinformer)

	if err != nil {
		log.Fatal(err)
	}

	defer connector.Close()

	for rows.Next() {
		err := rows.Scan(
			&idTrip,
			&shipName,
			&shipType,
			&tsOutOfSight,
			&shipFlag,
			&shipWidth,
			&shipLength,
			&grossTonnage,
			&netTonnage,
			&draftAft,
			&draftFwd,
			&agency,
			&lastPortOfCall,
			&portDestination,
		)

		if err != nil {
			log.Fatal(err)
		}

		tmpDict := map[string]string{
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

		result = append(result, tmpDict)
	}

	return result
}

// GetTodayShippedGoods todo doc
func GetTodayShippedGoods(idPortinformer string) []map[string]string {
	var idTrip, shipName, quantity sql.NullString
	var unit, goodsCategory, shipType, shipFlag, shipWidth, shipLength sql.NullString
	var grossTonnage, netTonnage, groupCategory, macroCategory sql.NullString

	result := []map[string]string{}

	connector := Connect()

	rows, err := odot.Query(connector, "shipped-goods", idPortinformer)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(
			&idTrip,
			&shipName,
			&quantity,
			&unit,
			&goodsCategory,
			&shipType,
			&shipFlag,
			&shipWidth,
			&shipLength,
			&grossTonnage,
			&netTonnage,
			&groupCategory,
			&macroCategory,
		)

		if err != nil {
			log.Fatal(err)
		}

		tmpDict := map[string]string{
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

		result = append(result, tmpDict)
	}

	return result
}

//GetTodayTrafficList todo doc
func GetTodayTrafficList(idPortinformer string) []map[string]string {
	var idTrip, shipName sql.NullString
	var numContainer, numPassengers, numCamion sql.NullString
	var numFurgoni, numRimorchi, numAuto, numMoto, numCamper, tons sql.NullString
	var numBus, numMinibus, mvntType, description sql.NullString
	var quay sql.NullString

	result := []map[string]string{}

	connector := Connect()

	rows, err := odot.Query(connector, "traffic-list", idPortinformer)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(
			&idTrip,
			&shipName,
			&numContainer,
			&numPassengers,
			&numCamion,
			&numFurgoni,
			&numRimorchi,
			&numAuto,
			&numMoto,
			&numCamper,
			&tons,
			&numBus,
			&numMinibus,
			&mvntType,
			&description,
			&quay,
		)

		if err != nil {
			log.Fatal(err)
		}

		tmpDict := map[string]string{
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

		result = append(result, tmpDict)
	}

	return result

}
