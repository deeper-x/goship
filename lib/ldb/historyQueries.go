package ldb

import (
	"database/sql"
	"fmt"
	"log"
)

//GetAllShippedGoods todo doc
func GetAllShippedGoods(idPortinformer string, idArrivalPrevisionState int, idDepartureState int) []map[string]string {
	var idTrip, shipName, length, width, grossTonnage sql.NullString
	var netTonnage, agencyArrival, shipFlag sql.NullString
	var operation, shipType, shipSubtype sql.NullString
	var arrPrevDraftAft, arrPrevDraftFwd sql.NullString
	var depDraftAft, depDraftFwd sql.NullString
	var agencyArrivPrev, agencyDep, destinationPort, destinationCountry sql.NullString
	var lastPortOfCallName sql.NullString
	var lastPortOfCallCountry, tsSighting, tsOutOfSight, shippedGoodsDetails sql.NullString
	var quantity, unit, shippedGoodsGroup, shippedGoodsMacro sql.NullString
	var quay, berth sql.NullString

	result := []map[string]string{}

	query := fmt.Sprintf(`SELECT id_control_unit_data, ships.ship_description AS ship_name,
		ships.length AS length, ships.width AS width, ships.gross_tonnage AS gross_tonnage,
		ships.net_tonnage AS net_tonnage,
		arrival_agency.description AS agency_arrival, countries.iso3 AS ship_flag,
		goods_mvmnt_type as operation,
		ship_types.type_description AS ship_type, ship_subtypes.description AS ship_subtype,
		maneuv_data_arriv_prev.draft_aft as arr_prev_draft_aft,
		maneuv_data_arriv_prev.draft_fwd as arr_prev_draft_fwd,
		maneuv_data_dep.draft_aft as dep_draft_aft,
		maneuv_data_dep.draft_fwd as dep_draft_fwd,
		maneuv_data_arriv_prev.agency_arriv_prev,
		maneuv_data_dep.agency_dep,
		destination_data.port_name AS destination_port,
		destination_data.port_country AS destination_country,
		lpc_data.port_name AS last_port_of_call_name,
		lpc_data.port_country AS last_port_of_call_country,
		data_avvistamento_nave.ts_avvistamento AS ts_sighting,
		data_fuori_dal_porto.ts_out_of_sight AS ts_out_of_sight,  
		goods_categories.description AS shipped_goods_details, 
		CASE WHEN quantity = '' THEN '0' ELSE quantity END, 
		unit, groups_categories.description AS shipped_goods_group,
		macro_categories.description AS shipped_goods_macro,
		quays.description AS quay, berths.description AS berth
		FROM shipped_goods
		INNER JOIN control_unit_data
		ON control_unit_data.id_control_unit_data = shipped_goods.fk_control_unit_data
		INNER JOIN goods_categories
		ON goods_categories.id_goods_category = shipped_goods.fk_goods_category
		INNER JOIN groups_categories
		ON groups_categories.id_group = goods_categories.fk_group_category
		INNER JOIN macro_categories
		ON groups_categories.fk_macro_category = macro_categories.id_macro_category
		INNER JOIN ships
		ON control_unit_data.fk_ship = ships.id_ship
		LEFT JOIN ship_types
		ON ships.fk_ship_type = ship_types.id_ship_type
		LEFT JOIN ship_subtypes
		ON ship_subtypes.id_ship_subtype = ships.fk_ship_subtype
		INNER JOIN data_avvistamento_nave
		ON control_unit_data.id_control_unit_data = data_avvistamento_nave.fk_control_unit_data
		LEFT JOIN data_fuori_dal_porto
		ON data_fuori_dal_porto.fk_control_unit_data = control_unit_data.id_control_unit_data
		INNER JOIN (
			SELECT id_agency, description
			FROM agencies
			WHERE fk_portinformer = %s
		) as arrival_agency
		ON arrival_agency.id_agency = data_avvistamento_nave.fk_agency
		INNER JOIN quays
		ON quays.id_quay = shipped_goods.fk_operation_quay
		INNER JOIN berths
		ON berths.id_berth = shipped_goods.fk_operation_berth
		LEFT JOIN countries
		ON ships.fk_country_flag = id_country 
		INNER JOIN (
			SELECT id_maneuvering, trips_logs.fk_state as trip_state,
			trips_logs.fk_control_unit_data as id_trip,
			draft_aft, draft_fwd, agencies.description AS agency_arriv_prev
			FROM trips_logs INNER JOIN maneuverings
			ON fk_maneuvering = id_maneuvering
			INNER JOIN agencies
			ON agencies.id_agency = trips_logs.fk_agency
			WHERE trips_logs.fk_portinformer = %s
			AND trips_logs.fk_state = %d
			GROUP BY id_maneuvering, agencies.description, trip_state, id_trip, draft_aft, draft_fwd
		) AS maneuv_data_arriv_prev
		ON maneuv_data_arriv_prev.id_trip = id_control_unit_data
		INNER JOIN (
			SELECT id_maneuvering, trips_logs.fk_state as trip_state,
			trips_logs.fk_control_unit_data as id_trip,
			draft_aft, draft_fwd, agencies.description AS agency_dep
			FROM trips_logs INNER JOIN maneuverings
			ON fk_maneuvering = id_maneuvering
			INNER JOIN agencies
			ON agencies.id_agency = trips_logs.fk_agency
			WHERE trips_logs.fk_portinformer = %s
			AND trips_logs.fk_state = %d
			GROUP BY id_maneuvering, agencies.description, trip_state, id_trip, draft_aft, draft_fwd
		) AS maneuv_data_dep
		ON maneuv_data_dep.id_trip = id_control_unit_data
		LEFT JOIN (
			SELECT id_shipping_details, ports.name AS port_name,
			ports.country as port_country
			FROM ports 
			INNER JOIN shipping_details
			ON shipping_details.fk_port_provenance = id_port
		) AS lpc_data
		ON lpc_data.id_shipping_details = control_unit_data.fk_shipping_details
		LEFT JOIN (
			SELECT id_shipping_details, ports.name AS port_name,
			ports.country as port_country
			FROM ports 
			INNER JOIN shipping_details
			ON shipping_details.fk_port_destination = id_port
		) AS destination_data
		ON destination_data.id_shipping_details = control_unit_data.fk_shipping_details
		WHERE control_unit_data.fk_portinformer = %s`, idPortinformer, idPortinformer, idArrivalPrevisionState, idPortinformer, idDepartureState, idPortinformer)

	connector := Connect()

	rows, err := connector.Query(query)

	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(
			&idTrip,
			&shipName,
			&length,
			&width,
			&grossTonnage,
			&netTonnage,
			&agencyArrival,
			&shipFlag,
			&operation,
			&shipType,
			&shipSubtype,
			&arrPrevDraftAft,
			&arrPrevDraftFwd,
			&depDraftAft,
			&depDraftFwd,
			&agencyArrivPrev,
			&agencyDep,
			&destinationPort,
			&destinationCountry,
			&lastPortOfCallName,
			&lastPortOfCallCountry,
			&tsSighting,
			&tsOutOfSight,
			&shippedGoodsDetails,
			&quantity,
			&unit,
			&shippedGoodsGroup,
			&shippedGoodsMacro,
			&quay,
			&berth,
		)

		if err != nil {
			log.Fatal(err)
		}

		tmpDict := map[string]string{
			"idTrip":                idTrip.String,
			"shipName":              shipName.String,
			"length":                length.String,
			"width":                 width.String,
			"grossTonnage":          grossTonnage.String,
			"netTonnage":            netTonnage.String,
			"agencyArrival":         agencyArrival.String,
			"shipFlag":              shipFlag.String,
			"operation":             operation.String,
			"shipType":              shipType.String,
			"shipSubtype":           shipSubtype.String,
			"arrPrevDraftAft":       arrPrevDraftAft.String,
			"arrPrevDraftFwd":       arrPrevDraftFwd.String,
			"depDraftAft":           depDraftAft.String,
			"depDraftFwd":           depDraftFwd.String,
			"agencyArrivPrev":       agencyArrivPrev.String,
			"agencyDep":             agencyDep.String,
			"destinationPort":       destinationPort.String,
			"destinationCountry":    destinationCountry.String,
			"lastPortOfCallName":    lastPortOfCallName.String,
			"lastPortOfCallCountry": lastPortOfCallCountry.String,
			"tsSighting":            tsSighting.String,
			"tsOutOfSight":          tsOutOfSight.String,
			"shippedGoodsDetails":   shippedGoodsDetails.String,
			"quantity":              quantity.String,
			"unit":                  unit.String,
			"shippedGoodsGroup":     shippedGoodsGroup.String,
			"shippedGoodsMacro":     shippedGoodsMacro.String,
			"quay":                  quay.String,
			"berth":                 berth.String,
		}

		result = append(result, tmpDict)

	}

	return result
}

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
