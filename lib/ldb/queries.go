package ldb

import (
	"database/sql"
	"fmt"
	"log"
)

//GetAllMoored todo doc
func GetAllMoored(idPortinformer string) []map[string]string {
	var idControlUnitData sql.NullString
	var shipName, mooringTime, mooringMooringTime, roadsteadMooringTime sql.NullString
	var storedMooringMooringTime, storedRoadsteadMooringTime sql.NullString
	var result []map[string]string

	connector := Connect()

	query := fmt.Sprintf(`SELECT id_control_unit_data, ship_description, 
						  data_ormeggio_nave.ts_fine_ormeggio,
						  data_da_ormeggio_a_ormeggio.ts_fine_ormeggio,
						  stored_data_da_ormeggio_a_ormeggio.ts_fine_ormeggio,
						  data_da_rada_a_ormeggio.ts_fine_ormeggio,
						  stored_data_da_rada_a_ormeggio.ts_fine_ormeggio 
						  FROM control_unit_data 
						  INNER JOIN ships
						  ON fk_ship = id_ship
						  LEFT JOIN data_ormeggio_nave
						  ON id_control_unit_data = data_ormeggio_nave.fk_control_unit_data
						  LEFT JOIN data_da_ormeggio_a_ormeggio
						  ON id_control_unit_data = data_da_ormeggio_a_ormeggio.fk_control_unit_data
						  LEFT JOIN stored_data_da_ormeggio_a_ormeggio
						  ON id_control_unit_data = stored_data_da_ormeggio_a_ormeggio.fk_control_unit_data
						  LEFT JOIN data_da_rada_a_ormeggio
						  ON id_control_unit_data = data_da_rada_a_ormeggio.fk_control_unit_data
						  LEFT JOIN stored_data_da_rada_a_ormeggio
						  ON id_control_unit_data = stored_data_da_rada_a_ormeggio.fk_control_unit_data
						  WHERE fk_ship_current_activity = 5 AND fk_portinformer = %s`, idPortinformer)

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
			&mooringMooringTime,
			&storedMooringMooringTime,
			&roadsteadMooringTime,
			&storedRoadsteadMooringTime)
		if err != nil {
			log.Fatal(err)
		}

		idControlUnitDataStr := idControlUnitData
		tmpDict := map[string]string{
			"id_trip":                       idControlUnitDataStr.String,
			"ship":                          shipName.String,
			"mooring_time":                  mooringTime.String,
			"mooring_mooring_time":          mooringMooringTime.String,
			"stored_mooring_mooring_time":   storedMooringMooringTime.String,
			"roadstead_mooring_time":        roadsteadMooringTime.String,
			"stored_roadstead_mooring_time": storedRoadsteadMooringTime.String}
		result = append(result, tmpDict)
	}

	return result
}
