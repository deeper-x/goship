package ldb

import (
	"database/sql"
	"fmt"
	"log"
)

// GetArrivalsRegister todo doc
func GetArrivalsRegister(idPortinformer string, idArrivalPrevision int, start string, stop string) []map[string]string {
	var idTrip, shipName, shipType, tsSighting, shipFlag, shipWidth, shipLength sql.NullString
	var grossTonnage, netTonnage, draftAft, draftFwd, agency, lastPortOfCall sql.NullString
	var portDestination, destinationQuayBerth, destinationRoadstead sql.NullString

	var result []map[string]string = []map[string]string{}

	connector := Connect()

	query := fmt.Sprintf(`SELECT id_control_unit_data AS id_trip, 
		ships.ship_description AS ship_name, 
		ship_types.type_acronym AS ship_type,  
		data_avvistamento_nave.ts_avvistamento AS ts_sighting, 
		countries.iso3 AS ship_flag,
		ships.width AS ship_width,
		ships.length AS ship_length,
		ships.gross_tonnage AS gross_tonnage,
		ships.net_tonnage AS net_tonnage,
		maneuverings.draft_aft AS draft_aft,
		maneuverings.draft_fwd AS draft_fwd,
		agencies.description AS agency,
		last_port_of_call.port_name||'('||last_port_of_call.port_country||')' AS last_port_of_call,
		port_destination.port_name||'('||port_destination.port_country||')' AS port_destination,
		quays.description AS destination_quay_berth,
		anchorage_points.description AS destination_roadstead
		FROM control_unit_data
		INNER JOIN data_avvistamento_nave
		ON data_avvistamento_nave.fk_control_unit_data = id_control_unit_data
		INNER JOIN ships
		ON control_unit_data.fk_ship = ships.id_ship
		INNER JOIN ship_types
		ON ships.fk_ship_type = ship_types.id_ship_type
		INNER JOIN countries
		ON ships.fk_country_flag = countries.id_country
		INNER JOIN maneuverings
		ON maneuverings.fk_control_unit_data = control_unit_data.id_control_unit_data
		AND maneuverings.fk_state = %d
		INNER JOIN agencies
		ON data_avvistamento_nave.fk_agency = agencies.id_agency
		INNER JOIN shipping_details
		ON control_unit_data.fk_shipping_details = shipping_details.id_shipping_details
		INNER JOIN (
			SELECT id_port, ports.name AS port_name, ports.country AS port_country
			FROM ports
		) AS last_port_of_call
		ON shipping_details.fk_port_provenance = last_port_of_call.id_port
		INNER JOIN (
			SELECT id_port, ports.name AS port_name, ports.country AS port_country
			FROM ports
		) AS port_destination
		ON shipping_details.fk_port_destination = port_destination.id_port
		LEFT JOIN quays
		ON maneuverings.fk_stop_quay = quays.id_quay
		AND maneuverings.fk_state = %d
		LEFT JOIN berths
		ON maneuverings.fk_stop_berth = berths.id_berth
		AND maneuverings.fk_state = %d
		LEFT JOIN anchorage_points
		ON maneuverings.fk_stop_anchorage_point = anchorage_points.id_anchorage_point
		AND maneuverings.fk_state = %d
		WHERE control_unit_data.fk_portinformer = %s
		AND LENGTH(ts_avvistamento) > 0
		AND ts_avvistamento::TIMESTAMP BETWEEN '%s' AND '%s'`, idArrivalPrevision, idArrivalPrevision, idArrivalPrevision, idArrivalPrevision, idPortinformer, start, stop)

	rows, err := connector.Query(query)

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
