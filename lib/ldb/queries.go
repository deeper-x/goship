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
