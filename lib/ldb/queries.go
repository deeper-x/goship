package ldb

import (
	"database/sql"
	"fmt"
	"log"
)

// GetAllRoadstead todo doc
func GetAllRoadstead(idPortinformer string) []map[string]string {
	var idControlUnitData sql.NullString
	var shipName, anchoringTime, currentActivity sql.NullString
	var anchoragePoint sql.NullString
	var result []map[string]string

	connector := Connect()

	query := fmt.Sprintf(`SELECT id_control_unit_data, ship_description, 
						  ts_last_ship_activity, 
						  ship_current_activities.description AS current_activity,
						  anchorage_points.description AS anchorage_point
						  FROM control_unit_data 
						  INNER JOIN ships
						  ON fk_ship = id_ship
						  INNER JOIN ship_current_activities
						  ON fk_ship_current_activity = id_activity
						  INNER JOIN latest_maneuverings
						  ON latest_maneuverings.fk_control_unit_data = id_control_unit_data
						  INNER JOIN anchorage_points
						  ON latest_maneuverings.fk_stop_anchorage_point = id_anchorage_point
						  WHERE fk_ship_current_activity = 2
						  AND is_active = true 
						  AND control_unit_data.fk_portinformer = %s`, idPortinformer)

	rows, err := connector.Query(query)

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
		)

		if err != nil {
			log.Fatal(err)
		}

		idControlUnitDataStr := idControlUnitData

		tmpDict := map[string]string{
			"id_trip":          idControlUnitDataStr.String,
			"ship":             shipName.String,
			"anchoring_time":   anchoringTime.String,
			"current_activity": currentActivity.String,
			"anchorage_point":  anchoragePoint.String,
		}

		result = append(result, tmpDict)
	}

	return result
}

//GetActiveArrivalPrevisions todo doc
func GetActiveArrivalPrevisions(idPortinformer string) []map[string]string {
	var idControlUnitData, shipName sql.NullString
	var tsArrivalPrevision, shipType sql.NullString
	var shipFlag, shipWidth, shipLength, grossTonnage sql.NullString
	var netTonnage, draftAft, draftFwd sql.NullString
	var agency sql.NullString
	var lastPortOfCall sql.NullString
	var destinationQuayBerth sql.NullString
	var destinationRoadstead sql.NullString

	var result []map[string]string = []map[string]string{}

	connector := Connect()

	query := fmt.Sprintf(`SELECT ship_description AS ship, ts_arrival_prevision,
			  ship_types.type_acronym AS ship_type,  
			  countries.iso3 AS ship_flag,
			  ships.width AS ship_width,
			  ships.length AS ship_length,
			  ships.gross_tonnage AS gross_tonnage,
			  ships.net_tonnage AS net_tonnage,
			  draft_aft, draft_fwd,
			  agencies.description AS agency,
			  last_port_of_call.port_name||'('||last_port_of_call.port_country||')' AS last_port_of_call,
			  quays.description AS destination_quay_berth,
			  anchorage_points.description AS destination_roadstead
			  FROM planned_arrivals
			  INNER JOIN ships
			  ON ships.id_ship = planned_arrivals.fk_ship
			  INNER JOIN ship_types
			  ON ships.fk_ship_type = ship_types.id_ship_type
			  INNER JOIN countries
			  ON ships.fk_country_flag = countries.id_country
			  INNER JOIN agencies
			  ON planned_arrivals.fk_agency = agencies.id_agency
			  INNER JOIN (
					SELECT id_port, ports.name AS port_name, ports.country AS port_country
					FROM ports
			  ) AS last_port_of_call
			  ON planned_arrivals.fk_last_port_of_call = last_port_of_call.id_port
			  LEFT JOIN quays
			  ON planned_arrivals.fk_stop_quay = quays.id_quay
			  LEFT JOIN berths
			  ON planned_arrivals.fk_stop_berth = berths.id_berth
			  LEFT JOIN anchorage_points
			  ON planned_arrivals.fk_stop_anchorage_point = anchorage_points.id_anchorage_point	
			  WHERE LENGTH(planned_arrivals.ts_arrival_prevision) > 0 
			  AND planned_arrivals.is_active = true
			  AND planned_arrivals.fk_portinformer = %s`, idPortinformer)

	rows, err := connector.Query(query)

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
		}

		result = append(result, tmpDict)
	}

	return result
}

//GetAllMoored todo doc
func GetAllMoored(idPortinformer string) []map[string]string {
	var idControlUnitData sql.NullString
	var shipName, mooringTime, currentActivity, quay sql.NullString
	var result []map[string]string

	connector := Connect()

	query := fmt.Sprintf(`SELECT id_control_unit_data, ship_description, 
						  ts_last_ship_activity, ship_current_activities.description AS current_activity, quays.description AS quay  
						  FROM control_unit_data 
						  INNER JOIN ships
						  ON fk_ship = id_ship
						  INNER JOIN ship_current_activities
						  ON fk_ship_current_activity = id_activity
						  INNER JOIN latest_maneuverings
						  ON latest_maneuverings.fk_control_unit_data = id_control_unit_data
						  INNER JOIN quays
						  ON latest_maneuverings.fk_stop_quay = id_quay
						  WHERE fk_ship_current_activity = 5
						  AND control_unit_data.is_active = true 
						  AND control_unit_data.fk_portinformer = %s`, idPortinformer)

	rows, err := connector.Query(query)

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
		}
		result = append(result, tmpDict)
	}

	return result
}

//GetTodayArrivals todo doc
func GetTodayArrivals(idPortinformer string, idArrivalPrevision string) []map[string]string {
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
		AND maneuverings.fk_state = %s
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
		AND maneuverings.fk_state = %s
		LEFT JOIN berths
		ON maneuverings.fk_stop_berth = berths.id_berth
		AND maneuverings.fk_state = %s
		LEFT JOIN anchorage_points
		ON maneuverings.fk_stop_anchorage_point = anchorage_points.id_anchorage_point
		AND maneuverings.fk_state = %s
		WHERE control_unit_data.fk_portinformer = %s
		AND LENGTH(ts_avvistamento) > 0
		AND ts_avvistamento::DATE = current_date`, idArrivalPrevision, idArrivalPrevision, idArrivalPrevision, idArrivalPrevision, idPortinformer)

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

// GetTodayDepartures todo description
func GetTodayDepartures(idPortinformer string, idDepartureState int) []map[string]string {
	var idTrip, shipName, shipType, tsOutOfSight, shipFlag, shipWidth sql.NullString
	var shipLength, grossTonnage sql.NullString
	var netTonnage, draftAft, draftFwd, agency, lastPortOfCall, portDestination sql.NullString

	var result []map[string]string = []map[string]string{}

	connector := Connect()

	query := fmt.Sprintf(`SELECT id_control_unit_data AS id_trip, 
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
			AND maneuverings.fk_state = %d
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
			WHERE control_unit_data.fk_portinformer = %s
			AND ts_out_of_sight IS NOT NULL
			AND ts_out_of_sight != 'None'
			AND LENGTH(ts_out_of_sight) > 0
			AND ts_out_of_sight::DATE = current_date`, idDepartureState, idPortinformer)

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
