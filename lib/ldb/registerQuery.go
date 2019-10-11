package ldb

import (
	"database/sql"
	"log"
)

// GetArrivalsRegister todo doc
func GetArrivalsRegister(idPortinformer string, idArrivalPrevision int, start string, stop string) []map[string]string {
	var idTrip, shipName, shipType, tsSighting, shipFlag, shipWidth, shipLength sql.NullString
	var grossTonnage, netTonnage, draftAft, draftFwd, agency, lastPortOfCall sql.NullString
	var portDestination, destinationQuayBerth, destinationRoadstead sql.NullString

	var result []map[string]string = []map[string]string{}

	connector := Connect()

	query := `SELECT id_control_unit_data AS id_trip, 
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
AND maneuverings.fk_state = $1
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
AND maneuverings.fk_state = $2
LEFT JOIN berths
ON maneuverings.fk_stop_berth = berths.id_berth
AND maneuverings.fk_state = $3
LEFT JOIN anchorage_points
ON maneuverings.fk_stop_anchorage_point = anchorage_points.id_anchorage_point
AND maneuverings.fk_state = $4
WHERE control_unit_data.fk_portinformer = $5
AND LENGTH(ts_avvistamento) > 0
AND ts_avvistamento::TIMESTAMP BETWEEN $6 AND $7`

	rows, err := connector.Query(query, idArrivalPrevision, idArrivalPrevision, idArrivalPrevision, idArrivalPrevision, idPortinformer, start, stop)

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

	query := `(SELECT id_control_unit_data, 
ship_description, type_description, ts_fine_ormeggio,
countries.name as country, width, length, gross_tonnage, net_tonnage 
FROM control_unit_data
INNER JOIN data_ormeggio_nave
ON fk_control_unit_data = id_control_unit_data
INNER JOIN ships
ON fk_ship = id_ship
INNER JOIN ship_types
ON fk_ship_type = id_ship_type
INNER JOIN countries
ON fk_country_flag = id_country
WHERE control_unit_data.fk_portinformer = $1 
AND ts_fine_ormeggio BETWEEN $2 AND $3)
UNION
(SELECT id_control_unit_data, 
	ship_description, type_description, ts_fine_ormeggio,
	countries.name as country, width, length, gross_tonnage, net_tonnage 
	FROM control_unit_data
	INNER JOIN data_da_ormeggio_a_ormeggio
	ON fk_control_unit_data = id_control_unit_data
	INNER JOIN ships
	ON fk_ship = id_ship
	INNER JOIN ship_types
	ON fk_ship_type = id_ship_type
	INNER JOIN countries
	ON fk_country_flag = id_country
	WHERE control_unit_data.fk_portinformer = $4 
	AND ts_fine_ormeggio BETWEEN $5 AND $6)
UNION
(
SELECT id_control_unit_data, 
ship_description, type_description, ts_fine_ormeggio,
countries.name as country, width, length, gross_tonnage, net_tonnage 
FROM control_unit_data
INNER JOIN data_da_rada_a_ormeggio
ON fk_control_unit_data = id_control_unit_data
INNER JOIN ships
ON fk_ship = id_ship
INNER JOIN ship_types
ON fk_ship_type = id_ship_type
INNER JOIN countries
ON fk_country_flag = id_country
WHERE control_unit_data.fk_portinformer = $7 
AND ts_fine_ormeggio BETWEEN $8 AND $9
)`

	rows, err := connector.Query(query, idPortinformer, start, stop, idPortinformer, start, stop, idPortinformer, start, stop)

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

	query := `(SELECT id_control_unit_data, 
ship_description, type_description, ts_anchor_drop,
countries.name as country, width, length, gross_tonnage, net_tonnage 
FROM control_unit_data
INNER JOIN data_arrivo_in_rada
ON fk_control_unit_data = id_control_unit_data
INNER JOIN ships
ON fk_ship = id_ship
INNER JOIN ship_types
ON fk_ship_type = id_ship_type
INNER JOIN countries
ON fk_country_flag = id_country
WHERE control_unit_data.fk_portinformer = $1 
AND ts_anchor_drop BETWEEN $2 AND $3)
UNION
(SELECT id_control_unit_data, 
	ship_description, type_description, ts_anchor_drop,
	countries.name as country, width, length, gross_tonnage, net_tonnage 
	FROM control_unit_data
	INNER JOIN data_da_ormeggio_a_rada
	ON fk_control_unit_data = id_control_unit_data
	INNER JOIN ships
	ON fk_ship = id_ship
	INNER JOIN ship_types
	ON fk_ship_type = id_ship_type
	INNER JOIN countries
	ON fk_country_flag = id_country
	WHERE control_unit_data.fk_portinformer = $4 
	AND ts_anchor_drop BETWEEN $5 AND $6)
UNION
(
SELECT id_control_unit_data, 
ship_description, type_description, ts_anchor_drop,
countries.name as country, width, length, gross_tonnage, net_tonnage 
FROM control_unit_data
INNER JOIN data_da_rada_a_rada
ON fk_control_unit_data = id_control_unit_data
INNER JOIN ships
ON fk_ship = id_ship
INNER JOIN ship_types
ON fk_ship_type = id_ship_type
INNER JOIN countries
ON fk_country_flag = id_country
WHERE control_unit_data.fk_portinformer = $7 
AND ts_anchor_drop BETWEEN $8 AND $9
)`

	rows, err := connector.Query(query, idPortinformer, start, stop, idPortinformer, start, stop, idPortinformer, start, stop)

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

	query := `SELECT id_control_unit_data AS id_trip, 
ships.ship_description AS ship_name, 
ship_types.type_acronym AS ship_type,  
data_fuori_dal_porto.ts_out_of_sight AS ts_out_of_sight, 
countries.iso3 AS ship_flag,
ships.width AS ship_width,
ships.length AS ship_length,
ships.gross_tonnage AS gross_tonnage,
ships.net_tonnage AS net_tonnage,
maneuverings.draft_aft AS draft_aft,
maneuverings.draft_fwd AS draft_fwd,
agencies.description AS agency,
last_port_of_call.port_name||'('||last_port_of_call.port_country||')' AS last_port_of_call,
port_destination.port_name||'('||port_destination.port_country||')' AS port_destination
FROM control_unit_data
INNER JOIN data_fuori_dal_porto
ON data_fuori_dal_porto.fk_control_unit_data = id_control_unit_data
INNER JOIN ships
ON control_unit_data.fk_ship = ships.id_ship
INNER JOIN ship_types
ON ships.fk_ship_type = ship_types.id_ship_type
INNER JOIN countries
ON ships.fk_country_flag = countries.id_country
INNER JOIN maneuverings
ON maneuverings.fk_control_unit_data = control_unit_data.id_control_unit_data
AND maneuverings.fk_state = $1
INNER JOIN agencies
ON data_fuori_dal_porto.fk_agency = agencies.id_agency
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
WHERE control_unit_data.fk_portinformer = $2
AND ts_out_of_sight IS NOT NULL
AND ts_out_of_sight != 'None'
AND LENGTH(ts_out_of_sight) > 0
AND ts_out_of_sight::TIMESTAMP BETWEEN $3 AND $4`

	rows, err := connector.Query(query, idDepartureState, idPortinformer, start, stop)

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
