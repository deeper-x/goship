package ldb

import (
	"database/sql"
	"log"

	"github.com/deeper-x/goship/conf"
)

// GetArrivalsRegister todo doc
func (r repository) GetArrivalsRegister(idPortinformer string, idArrivalPrevision int, start string, stop string) []map[string]string {
	var idTrip, shipName, shipType, tsSighting, shipFlag, shipWidth, shipLength sql.NullString
	var grossTonnage, netTonnage, draftAft, draftFwd, agency, lastPortOfCall sql.NullString
	var portDestination, destinationQuayBerth, destinationRoadstead sql.NullString

	var result = []map[string]string{}

	mapper.GenResource(conf.PRegisterSQL)
	rows, err := mapper.resource.Query(r.db, "arrivals-register", idArrivalPrevision, idArrivalPrevision, idArrivalPrevision, idArrivalPrevision, idPortinformer, start, stop)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

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

// GetShiftingsRegister todo doc
func (r repository) GetShiftingsRegister(idPortinformer string, start string, stop string) []map[string]string {
	var idTrip, tsSighting, imo, ship sql.NullString
	var shipType, iso3, fromQuay, toQuay, fromAnch, toAnch, agency sql.NullString

	var result = []map[string]string{}

	mapper.GenResource(conf.PRegisterSQL)
	rows, err := mapper.resource.Query(r.db, "shiftings-register", idPortinformer, start, stop)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&idTrip,
			&tsSighting,
			&imo,
			&ship,
			&shipType,
			&iso3,
			&fromQuay,
			&toQuay,
			&fromAnch,
			&toAnch,
			&agency,
		)
		if err != nil {
			log.Fatal(err)
		}

		tmpDict := map[string]string{
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
			"agency":      agency.String,
		}

		result = append(result, tmpDict)
	}

	return result
}

// GetMooredRegister todo doc
func (r repository) GetMooredRegister(idPortinformer string, start string, stop string) []map[string]string {
	var idTrip, shipName, shipType, tsMooring, shipFlag, shipWidth sql.NullString
	var shipLength, grossTonnage, stopQuay sql.NullString
	var netTonnage, agency sql.NullString

	var result = []map[string]string{}

	mapper.GenResource(conf.PRegisterSQL)

	rows, err := mapper.resource.Query(r.db, "moored-register", idPortinformer, start, stop)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&idTrip,
			&shipName,
			&shipType,
			&tsMooring,
			&shipFlag,
			&shipWidth,
			&shipLength,
			&grossTonnage,
			&netTonnage,
			&stopQuay,
			&agency,
		)

		if err != nil {
			log.Fatal(err)
		}

		tmpDict := map[string]string{
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
			"stop_quay":     stopQuay.String,
		}

		result = append(result, tmpDict)
	}

	return result

}

// GetRoadsteadRegister todo doc
func (r repository) GetRoadsteadRegister(idPortinformer string, start string, stop string) []map[string]string {
	var idTrip, shipName, shipType, tsAnchoring, shipFlag, shipWidth sql.NullString
	var shipLength, grossTonnage, roadstead sql.NullString
	var netTonnage, agency sql.NullString

	var result = []map[string]string{}

	mapper.GenResource(conf.PRegisterSQL)
	rows, err := mapper.resource.Query(r.db, "roadstead-register", idPortinformer, start, stop)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&idTrip,
			&shipName,
			&shipType,
			&tsAnchoring,
			&shipFlag,
			&shipWidth,
			&shipLength,
			&grossTonnage,
			&netTonnage,
			&roadstead,
			&agency,
		)

		if err != nil {
			log.Fatal(err)
		}

		tmpDict := map[string]string{
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
			"roadstead":     roadstead.String,
		}

		result = append(result, tmpDict)
	}

	return result

}

// GetShippedGoodsRegister todo description
func (r repository) GetShippedGoodsRegister(idPortinformer string, start string, stop string) []map[string]string {
	var idTrip, shipName, quantity sql.NullString
	var unit, goodsCategory, shipType, shipFlag, shipWidth, shipLength sql.NullString
	var grossTonnage, netTonnage, groupCategory, macroCategory sql.NullString

	result := []map[string]string{}

	mapper.GenResource(conf.PRegisterSQL)
	rows, err := mapper.resource.Query(r.db, "shipped-goods-register", idPortinformer, start, stop)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

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

// GetDeparturesRegister todo description
func (r repository) GetDeparturesRegister(idPortinformer string, idDepartureState int, start string, stop string) []map[string]string {
	var idTrip, shipName, shipType, tsOutOfSight, shipFlag, shipWidth sql.NullString
	var shipLength, grossTonnage sql.NullString
	var netTonnage, draftAft, draftFwd, agency, lastPortOfCall, portDestination sql.NullString

	var result = []map[string]string{}

	mapper.GenResource(conf.PRegisterSQL)
	rows, err := mapper.resource.Query(r.db, "departures-register", idDepartureState, idPortinformer, start, stop)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

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

//GetRegisterTrafficList todo doc
func (r repository) GetRegisterTrafficList(idPortinformer string, start string, stop string) []map[string]string {
	var idTrip, shipName, tsSighting sql.NullString
	var numContainer, numPassengers, numCamion sql.NullString
	var numFurgoni, numRimorchi, numAuto, numMoto, numCamper, tons sql.NullString
	var numBus, numMinibus, mvntType, description sql.NullString
	var quay sql.NullString

	result := []map[string]string{}

	mapper.GenResource(conf.PRegisterSQL)
	rows, err := mapper.resource.Query(r.db, "traffic-list-register", idPortinformer, start, stop)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&idTrip,
			&shipName,
			&tsSighting,
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

		result = append(result, tmpDict)
	}

	return result

}
