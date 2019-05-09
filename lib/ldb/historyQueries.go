package ldb

import (
	"database/sql"
	"fmt"
	"log"
)

//GetAllArrivalsArchive todo doc
func GetAllArrivalsArchive(idPortinformer string, idState int) []map[string]string {
	var result = []map[string]string{}

	var idTrip, shipName, shipType sql.NullString
	var tsSighting, shipFlag, shipWidth, shipLength sql.NullString
	var grossTonnage, netTonnage, draftAft sql.NullString
	var draftFwd, agency, lastPortOfCall, portDestination, destinationQuayBerth sql.NullString

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
		quays.description AS destination_quay_berth
		FROM control_unit_data
		INNER JOIN data_avvistamento_nave
		ON data_avvistamento_nave.fk_control_unit_data = id_control_unit_data
		INNER JOIN ships
		ON control_unit_data.fk_ship = ships.id_ship
		INNER JOIN ship_types
		ON ships.fk_ship_type = ship_types.id_ship_type
		LEFT JOIN countries
		ON ships.fk_country_flag = countries.id_country
		INNER JOIN maneuverings
		ON maneuverings.fk_control_unit_data = control_unit_data.id_control_unit_data
		AND maneuverings.fk_state = %d
		LEFT JOIN agencies
		ON data_avvistamento_nave.fk_agency = agencies.id_agency
		INNER JOIN shipping_details
		ON control_unit_data.fk_shipping_details = shipping_details.id_shipping_details
		LEFT JOIN (
			SELECT id_port, ports.name AS port_name, ports.country AS port_country
			FROM ports
		) AS last_port_of_call
		ON shipping_details.fk_port_provenance = last_port_of_call.id_port
		LEFT JOIN (
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
		WHERE control_unit_data.fk_portinformer = %s
		AND LENGTH(ts_avvistamento) > 0`, idState, idState, idState, idPortinformer)

	connector := Connect()

	rows, err := connector.Query(query)

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
		)

		if err != nil {
			log.Fatal(err)
		}

		tmpDict := map[string]string{
			"idTrip":               idTrip.String,
			"shipName":             shipName.String,
			"shipType":             shipType.String,
			"tsSighting":           tsSighting.String,
			"shipFlag":             shipFlag.String,
			"shipWidth":            shipWidth.String,
			"shipLength":           shipLength.String,
			"grossTonnage":         grossTonnage.String,
			"netTonnage":           netTonnage.String,
			"draftAft":             draftAft.String,
			"draftFwd":             draftAft.String,
			"agency":               agency.String,
			"lastPortOfCall":       lastPortOfCall.String,
			"portDestination":      portDestination.String,
			"destinationQuayBerth": destinationQuayBerth.String,
		}

		result = append(result, tmpDict)
	}

	return result
}
