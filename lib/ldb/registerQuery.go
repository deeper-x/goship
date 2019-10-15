package ldb

import (
	"database/sql"
	"log"

	"github.com/deeper-x/goship/conf"
)

// GetArrivalsRegister todo doc
func GetArrivalsRegister(idPortinformer string, idArrivalPrevision int, start string, stop string) []map[string]string {
	var idTrip, shipName, shipType, tsSighting, shipFlag, shipWidth, shipLength sql.NullString
	var grossTonnage, netTonnage, draftAft, draftFwd, agency, lastPortOfCall sql.NullString
	var portDestination, destinationQuayBerth, destinationRoadstead sql.NullString

	var result []map[string]string = []map[string]string{}

	connector := Connect()

	mapper.GenResource(conf.PRegisterSQL)
	rows, err := mapper.resource.Query(connector, "arrivals-register", idArrivalPrevision, idArrivalPrevision, idArrivalPrevision, idArrivalPrevision, idPortinformer, start, stop)

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

// GetMooredRegister todo doc
func GetMooredRegister(idPortinformer string, start string, stop string) []map[string]string {
	var idTrip, shipName, shipType, tsMooring, shipFlag, shipWidth sql.NullString
	var shipLength, grossTonnage sql.NullString
	var netTonnage, agency sql.NullString

	var result []map[string]string = []map[string]string{}

	connector := Connect()

	mapper.GenResource(conf.PRegisterSQL)
	rows, err := mapper.resource.Query(connector, "moored-register", idPortinformer, start, stop, idPortinformer, start, stop, idPortinformer, start, stop)

	if err != nil {
		log.Fatal(err)
	}

	defer connector.Close()

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
		}

		result = append(result, tmpDict)
	}

	return result

}

// GetRoadsteadRegister todo doc
func GetRoadsteadRegister(idPortinformer string, start string, stop string) []map[string]string {
	var idTrip, shipName, shipType, tsAnchoring, shipFlag, shipWidth sql.NullString
	var shipLength, grossTonnage sql.NullString
	var netTonnage, agency sql.NullString

	var result []map[string]string = []map[string]string{}

	connector := Connect()

	mapper.GenResource(conf.PRegisterSQL)
	rows, err := mapper.resource.Query(connector, "roadstead-register", idPortinformer, start, stop, idPortinformer, start, stop, idPortinformer, start, stop)

	if err != nil {
		log.Fatal(err)
	}

	defer connector.Close()

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
		}

		result = append(result, tmpDict)
	}

	return result

}

// GetDeparturesRegister todo description
func GetDeparturesRegister(idPortinformer string, idDepartureState int, start string, stop string) []map[string]string {
	var idTrip, shipName, shipType, tsOutOfSight, shipFlag, shipWidth sql.NullString
	var shipLength, grossTonnage sql.NullString
	var netTonnage, draftAft, draftFwd, agency, lastPortOfCall, portDestination sql.NullString

	var result []map[string]string = []map[string]string{}

	connector := Connect()

	mapper.GenResource(conf.PRegisterSQL)
	rows, err := mapper.resource.Query(connector, "departures-register", idDepartureState, idPortinformer, start, stop)

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
