package ldb

import (
	"database/sql"
	"fmt"
	"log"
)

// GetAllRoadstead todo doc
func GetAllRoadstead(idPortinformer string) []map[string]string {
	var idControlUnitData sql.NullString
	var shipName, anchoringTime, currentActivity, shippedGoods sql.NullString
	var shipType, iso3, grossTonnage, anchoragePoint, length, width, agency sql.NullString
	var result []map[string]string = []map[string]string{}

	connector := Connect()

	query := fmt.Sprintf(`SELECT id_control_unit_data, ship_description, 
						  ts_last_ship_activity, 
						  ship_current_activities.description AS current_activity,
						  anchorage_points.description AS anchorage_point,
						  type_acronym as ship_type, iso3, gross_tonnage, 
						  ships.length, ships.width,
						  agencies.description as agency,
						  shipped_goods_data.shipped_goods_row AS shipped_goods_data
						  FROM control_unit_data 
						  INNER JOIN ships
						  ON fk_ship = id_ship
						  INNER JOIN ship_current_activities
						  ON fk_ship_current_activity = id_activity
						  INNER JOIN latest_maneuverings
						  ON latest_maneuverings.fk_control_unit_data = id_control_unit_data
						  INNER JOIN anchorage_points
						  ON latest_maneuverings.fk_stop_anchorage_point = id_anchorage_point
						  INNER JOIN ship_types
        				  ON ships.fk_ship_type = ship_types.id_ship_type
						  INNER JOIN countries
						  ON countries.id_country = ships.fk_country_flag
						  INNER JOIN (  
							SELECT fk_control_unit_data, MAX(ts_main_event_field_val) AS max_time, fk_agency
							FROM trips_logs
							WHERE fk_portinformer = %s
							GROUP BY fk_control_unit_data, fk_portinformer, fk_agency
							) AS RES
						  ON id_control_unit_data = RES.fk_control_unit_data
						  INNER JOIN agencies
						  ON RES.fk_agency = agencies.id_agency
						  LEFT JOIN (
							SELECT fk_control_unit_data, string_agg(goods_mvmnt_type||':'||goods_categories.description::TEXT||'-'||groups_categories.description||' ('||quantity||' '||unit::TEXT||')', ', ') AS shipped_goods_row
							FROM shipped_goods
							INNER JOIN goods_categories
							ON goods_categories.id_goods_category = shipped_goods.fk_goods_category
							INNER JOIN groups_categories
							ON groups_categories.id_group = goods_categories.fk_group_category
							GROUP BY fk_control_unit_data        
						  ) as shipped_goods_data
						  ON shipped_goods_data.fk_control_unit_data = control_unit_data.id_control_unit_data
						  WHERE fk_ship_current_activity = 2
						  AND is_active = true 
						  AND control_unit_data.fk_portinformer = %s`, idPortinformer, idPortinformer)

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

//GetTodayArrivalPrevisions todo doc
func GetTodayArrivalPrevisions(idPortinformer string) []map[string]string {
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
			  anchorage_points.description AS destination_roadstead,
			  cargo_on_board
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

//GetTodayShiftingPrevisions todo doc
func GetTodayShiftingPrevisions(idPortinformer string) []map[string]string {
	var ship, tsShiftingPrevision, shipType, shipFlag, shipWidth sql.NullString
	var shipLength, grossTonnage, netTonnage, draftAft, draftFwd sql.NullString
	var agency, destinationPort, startingQuayBerth, startingRoadstead, cargoOnBoard sql.NullString

	var result []map[string]string = []map[string]string{}

	query := fmt.Sprintf(`SELECT ship_description AS ship, ts_shifting_prevision,
			ship_types.type_acronym AS ship_type,  
			countries.iso3 AS ship_flag,
			ships.width AS ship_width,
			ships.length AS ship_length,
			ships.gross_tonnage AS gross_tonnage,
			ships.net_tonnage AS net_tonnage,
			planned_shiftings.draft_aft, planned_shiftings.draft_fwd,
			agencies.description AS agency,
			start_quay.description AS starting_quay_berth,
			start_anchorage_point.description AS starting_roadstead,
			stop_quay.description AS stop_quay_berth,
			stop_anchorage_point.description AS stop_roadstead,
			planned_shiftings.cargo_on_board
			FROM planned_shiftings
			INNER JOIN planned_arrivals
			ON planned_shiftings.fk_planned_arrival = planned_arrivals.id_planned_arrival
			INNER JOIN ships
			ON ships.id_ship = planned_arrivals.fk_ship
			INNER JOIN ship_types
			ON ships.fk_ship_type = ship_types.id_ship_type
			INNER JOIN countries
			ON ships.fk_country_flag = countries.id_country
			INNER JOIN agencies
			ON planned_shiftings.fk_agency = agencies.id_agency
			LEFT JOIN (
				select id_quay, description from quays
			) as start_quay
			ON planned_shiftings.fk_start_quay = start_quay.id_quay
			LEFT JOIN (
				select id_quay, description from quays
			) as stop_quay
			ON planned_shiftings.fk_stop_quay = stop_quay.id_quay
			LEFT JOIN (
				select id_berth, description from berths
			) as start_berth
			ON planned_shiftings.fk_start_berth = start_berth.id_berth
			LEFT JOIN (
				select id_berth, description from berths
			) as stop_berth
			ON planned_shiftings.fk_stop_berth = stop_berth.id_berth
			LEFT JOIN (
				select id_anchorage_point, description from anchorage_points
			) as start_anchorage_point
			ON planned_shiftings.fk_start_anchorage_point = start_anchorage_point.id_anchorage_point
			LEFT JOIN (
				select id_anchorage_point, description from anchorage_points
			) as stop_anchorage_point
			ON planned_shiftings.fk_stop_anchorage_point = stop_anchorage_point.id_anchorage_point	
			WHERE LENGTH(planned_shiftings.ts_shifting_prevision) > 0 
			AND planned_shiftings.is_active = true
			AND planned_shiftings.fk_portinformer = %s`, idPortinformer)

	connector := Connect()

	rows, err := connector.Query(query)

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
			&destinationPort,
			&startingQuayBerth,
			&startingRoadstead,
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

//GetTodayDeparturePrevisions todo doc
func GetTodayDeparturePrevisions(idPortinformer string) []map[string]string {
	var ship, tsDeparturePrevision, shipType, shipFlag, shipWidth sql.NullString
	var shipLength, grossTonnage, netTonnage, draftAft, draftFwd sql.NullString
	var agency, destinationPort, startingQuayBerth, startingRoadstead, cargoOnBoard sql.NullString

	var result []map[string]string = []map[string]string{}

	query := fmt.Sprintf(`SELECT ship_description AS ship, ts_departure_prevision,
			  ship_types.type_acronym AS ship_type,  
			  countries.iso3 AS ship_flag,
			  ships.width AS ship_width,
			  ships.length AS ship_length,
			  ships.gross_tonnage AS gross_tonnage,
			  ships.net_tonnage AS net_tonnage,
			  planned_departures.draft_aft, planned_departures.draft_fwd,
			  agencies.description AS agency,
			  destination_port.port_name||'('||destination_port.port_country||')' AS destination_port,
			  quays.description AS starting_quay_berth,
			  anchorage_points.description AS starting_roadstead,
			  planned_departures.cargo_on_board
			  FROM planned_departures
			  INNER JOIN planned_arrivals
			  ON planned_departures.fk_planned_arrival = planned_arrivals.id_planned_arrival
			  INNER JOIN ships
			  ON ships.id_ship = planned_arrivals.fk_ship
			  INNER JOIN ship_types
			  ON ships.fk_ship_type = ship_types.id_ship_type
			  INNER JOIN countries
			  ON ships.fk_country_flag = countries.id_country
			  INNER JOIN agencies
			  ON planned_departures.fk_agency = agencies.id_agency
			  INNER JOIN (
					SELECT id_port, ports.name AS port_name, ports.country AS port_country
					FROM ports
			  ) AS destination_port
			  ON planned_departures.fk_destination_port = destination_port.id_port
			  LEFT JOIN quays
			  ON planned_departures.fk_start_quay = quays.id_quay
			  LEFT JOIN berths
			  ON planned_departures.fk_start_berth = berths.id_berth
			  LEFT JOIN anchorage_points
			  ON planned_departures.fk_start_anchorage_point = anchorage_points.id_anchorage_point	
			  WHERE LENGTH(planned_departures.ts_departure_prevision) > 0 
			  AND planned_departures.is_active = true
			  AND planned_departures.fk_portinformer = %s`, idPortinformer)

	connector := Connect()

	rows, err := connector.Query(query)

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

	query := fmt.Sprintf(`SELECT id_control_unit_data, ship_description, 
						  ts_last_ship_activity, 
						  ship_current_activities.description AS current_activity, 
						  quays.description AS quay,
						  shipped_goods_data.shipped_goods_row AS shipped_goods_data,
						  iso3, gross_tonnage, ships.length, ships.width, type_acronym,
						  agencies.description AS agency  
						  FROM control_unit_data 
						  INNER JOIN ships
						  ON fk_ship = id_ship
						  INNER JOIN ship_current_activities
						  ON fk_ship_current_activity = id_activity
						  INNER JOIN latest_maneuverings
						  ON latest_maneuverings.fk_control_unit_data = id_control_unit_data
						  INNER JOIN quays
						  ON latest_maneuverings.fk_stop_quay = id_quay
						  LEFT JOIN (
							SELECT fk_control_unit_data, string_agg(goods_mvmnt_type||':'||goods_categories.description::TEXT||'-'||groups_categories.description||' ('||quantity||' '||unit::TEXT||')', ', ') AS shipped_goods_row
							FROM shipped_goods
							INNER JOIN goods_categories
							ON goods_categories.id_goods_category = shipped_goods.fk_goods_category
							INNER JOIN groups_categories
							ON groups_categories.id_group = goods_categories.fk_group_category
							GROUP BY fk_control_unit_data        
						  ) as shipped_goods_data
						  ON shipped_goods_data.fk_control_unit_data = control_unit_data.id_control_unit_data
						  INNER JOIN (  
							SELECT fk_control_unit_data, MAX(ts_main_event_field_val) AS max_time, fk_agency
							FROM trips_logs
							WHERE fk_portinformer = %s
							GROUP BY fk_control_unit_data, fk_portinformer, fk_agency
							) AS RES
						  ON id_control_unit_data = RES.fk_control_unit_data 
						  INNER JOIN countries
						  ON countries.id_country = ships.fk_country_flag
						  INNER JOIN ship_types
						  ON ships.fk_ship_type = ship_types.id_ship_type
						  INNER JOIN agencies
						  ON RES.fk_agency = agencies.id_agency 
						  WHERE fk_ship_current_activity = 5
						  AND control_unit_data.is_active = true 
						  AND control_unit_data.fk_portinformer = %s`, idPortinformer, idPortinformer)

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

// GetTodayShippedGoods todo doc
func GetTodayShippedGoods(idPortinformer string) []map[string]string {
	var idTrip, shipName, quantity sql.NullString
	var unit, goodsCategory, shipType, shipFlag, shipWidth, shipLength sql.NullString
	var grossTonnage, netTonnage, groupCategory, macroCategory sql.NullString

	result := []map[string]string{}

	query := fmt.Sprintf(`SELECT
		fk_control_unit_data AS id_trip,
		ships.ship_description AS ship_name, 
		CASE WHEN quantity = '' THEN '0' ELSE quantity END, 
		unit, 
		goods_categories.description AS goods_category,
		ship_types.type_acronym AS ship_type,  
		countries.iso3 AS ship_flag,
		ships.width AS ship_width,
		ships.length AS ship_length,
		ships.gross_tonnage AS gross_tonnage,
		ships.net_tonnage AS net_tonnage,
		groups_categories.description AS group_category,
		macro_categories.description AS macro_category                 
		FROM shipped_goods INNER JOIN control_unit_data
		ON fk_control_unit_data = id_control_unit_data
		INNER JOIN goods_categories
		ON fk_goods_category = id_goods_category
		INNER JOIN ships
		ON control_unit_data.fk_ship = id_ship
		INNER JOIN countries
		ON ships.fk_country_flag = id_country
		INNER JOIN ship_types
		ON ships.fk_ship_type = ship_types.id_ship_type
		INNER JOIN groups_categories
		ON goods_categories.fk_group_category = groups_categories.id_group
		INNER JOIN macro_categories
		ON groups_categories.fk_macro_category = macro_categories.id_macro_category     
		WHERE control_unit_data.fk_portinformer = %s
		AND control_unit_data.is_active = true`, idPortinformer)

	connector := Connect()

	rows, err := connector.Query(query)

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

	query := fmt.Sprintf(`SELECT
		control_unit_data.id_control_unit_data AS id_trip,
		ships.ship_description AS ship_name, 
		num_container, num_passengers, num_camion, 
		num_furgoni, num_rimorchi, num_auto, num_moto, num_camper, tons,
		num_bus, num_minibus, traffic_list_mvnt_type, traffic_list_categories.description,
		quays.description AS quay
		FROM traffic_list INNER JOIN control_unit_data
		ON fk_control_unit_data = id_control_unit_data
		INNER JOIN traffic_list_categories
		ON fk_traffic_list_category = id_traffic_list_category
		INNER JOIN ships
		ON control_unit_data.fk_ship = id_ship
		INNER JOIN maneuverings
		ON maneuverings.fk_control_unit_data = control_unit_data.id_control_unit_data
		AND maneuverings.fk_state = 17
		INNER JOIN quays
		ON maneuverings.fk_stop_quay = quays.id_quay
		WHERE control_unit_data.fk_portinformer = %s
		AND control_unit_data.is_active = true`, idPortinformer)

	connector := Connect()

	rows, err := connector.Query(query)

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
