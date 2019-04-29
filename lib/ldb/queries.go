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

//GetAllArrivalPrevisions todo doc
func GetAllArrivalPrevisions(idPortinformer string) []map[string]string {
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
			  WHERE planned_arrivals.fk_portinformer = %s`, idPortinformer)

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
