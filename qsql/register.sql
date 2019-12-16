--name:arrivals-register
SELECT id_control_unit_data AS id_trip, 
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
AND ts_avvistamento::TIMESTAMP BETWEEN $6 AND $7

--name: shiftings-register
SELECT id_control_unit_data, ts_main_event_field_val, 
imo, ship_description AS ship_name, 
type_acronym AS ship_type, iso3 AS country, 
start_quay.description||'/'||start_berth.description as FROM_QUAY,
stop_quay.description||'/'||stop_berth.description as TO_QUAY,
start_anchorage.description as FROM_ANCH,
stop_anchorage.description as TO_ANCH     
FROM control_unit_data 
INNER JOIN trips_logs
ON id_control_unit_data = trips_logs.fk_control_unit_data
INNER JOIN ships
ON control_unit_data.fk_ship = id_ship
INNER JOIN ship_types
ON id_ship_type = fk_ship_type
INNER JOIN countries
ON ships.fk_country_flag = id_country
INNER JOIN maneuverings
ON id_maneuvering = trips_logs.fk_maneuvering
INNER JOIN (
    SELECT id_quay, description 
    FROM quays 
) AS start_quay
ON start_quay.id_quay = maneuverings.fk_start_quay
INNER JOIN (
    SELECT id_quay, description 
    FROM quays 
) AS stop_quay
ON stop_quay.id_quay = maneuverings.fk_stop_quay
INNER JOIN (
    SELECT id_berth, description
    FROM berths
) AS start_berth
ON start_berth.id_berth = maneuverings.fk_start_berth
INNER JOIN (
    SELECT id_berth, description
    FROM berths
) AS stop_berth
ON stop_berth.id_berth = maneuverings.fk_stop_berth
INNER JOIN (
    SELECT id_anchorage_point, description
    FROM anchorage_points
) AS start_anchorage
ON start_anchorage.id_anchorage_point = maneuverings.fk_start_anchorage_point
INNER JOIN (
    SELECT id_anchorage_point, description
    FROM anchorage_points
) AS stop_anchorage
ON stop_anchorage.id_anchorage_point = maneuverings.fk_stop_anchorage_point
WHERE control_unit_data.fk_portinformer = $1
AND trips_logs.fk_state IN (18, 19, 20, 21, 22, 27)
AND ts_main_event_field_val 
BETWEEN $2 AND $3 
ORDER BY ts_main_event_field_val 



--name: moored-register
(SELECT id_control_unit_data, ship_description, type_description, ts_fine_ormeggio,
countries.name as country, width, length, gross_tonnage, net_tonnage, position_data.quay AS stop_quay 
FROM control_unit_data
INNER JOIN data_ormeggio_nave
ON fk_control_unit_data = id_control_unit_data
INNER JOIN ships
ON fk_ship = id_ship
INNER JOIN ship_types
ON fk_ship_type = id_ship_type
INNER JOIN countries
ON fk_country_flag = id_country
LEFT JOIN (
	SELECT trips_logs.fk_control_unit_data AS id_trip_data, quays.description AS quay
	FROM trips_logs 
	INNER JOIN maneuverings
	ON trips_logs.fk_maneuvering = id_maneuvering
	INNER JOIN quays 
	ON maneuverings.fk_stop_quay = id_quay
	INNER JOIN data_ormeggio_nave
	ON trips_logs.fk_control_unit_data = data_ormeggio_nave.fk_control_unit_data 
	WHERE ts_main_event_field_val BETWEEN $2 AND $3
	AND trips_logs.data_table_id::INTEGER = data_ormeggio_nave.id_data_ormeggio_nave
	AND trips_logs.fk_portinformer = $1
	AND fk_maneuvering IS NOT NULL
	AND maneuverings.fk_stop_quay != 0
	ORDER BY ts_main_event_field_val ASC
) AS position_data
ON id_control_unit_data = position_data.id_trip_data
WHERE control_unit_data.fk_portinformer = $1 
AND ts_fine_ormeggio BETWEEN $2 AND $3)
UNION
(SELECT id_control_unit_data, 
	ship_description, type_description, ts_fine_ormeggio,
	countries.name as country, width, length, gross_tonnage, net_tonnage, position_data.quay AS stop_quay
	FROM control_unit_data
	INNER JOIN data_da_ormeggio_a_ormeggio
	ON fk_control_unit_data = id_control_unit_data
	INNER JOIN ships
	ON fk_ship = id_ship
	INNER JOIN ship_types
	ON fk_ship_type = id_ship_type
	INNER JOIN countries
	ON fk_country_flag = id_country
	LEFT JOIN (
		SELECT trips_logs.fk_control_unit_data AS id_trip_data, quays.description AS quay
		FROM trips_logs 
		INNER JOIN maneuverings
		ON trips_logs.fk_maneuvering = id_maneuvering
		INNER JOIN quays 
		ON maneuverings.fk_stop_quay = id_quay
		INNER JOIN data_da_ormeggio_a_ormeggio
		ON trips_logs.fk_control_unit_data = data_da_ormeggio_a_ormeggio.fk_control_unit_data 
		WHERE ts_main_event_field_val BETWEEN $2 AND $3
		AND trips_logs.data_table_id::INTEGER = data_da_ormeggio_a_ormeggio.id_data_da_ormeggio_a_ormeggio
		AND trips_logs.fk_portinformer = $1
		AND fk_maneuvering IS NOT NULL
		AND maneuverings.fk_stop_quay != 0
		ORDER BY ts_main_event_field_val ASC
	) AS position_data
	ON id_control_unit_data = position_data.id_trip_data
	WHERE control_unit_data.fk_portinformer = $1 
	AND ts_fine_ormeggio BETWEEN $2 AND $3)
UNION
(
SELECT id_control_unit_data, 
ship_description, type_description, ts_fine_ormeggio,
countries.name as country, width, length, gross_tonnage, net_tonnage, position_data.quay AS stop_quay
FROM control_unit_data
INNER JOIN data_da_rada_a_ormeggio
ON fk_control_unit_data = id_control_unit_data
INNER JOIN ships
ON fk_ship = id_ship
INNER JOIN ship_types
ON fk_ship_type = id_ship_type
INNER JOIN countries
ON fk_country_flag = id_country
LEFT JOIN (
	SELECT trips_logs.fk_control_unit_data AS id_trip_data, quays.description AS quay
	FROM trips_logs 
	INNER JOIN maneuverings
	ON trips_logs.fk_maneuvering = id_maneuvering
	INNER JOIN quays 
	ON maneuverings.fk_stop_quay = id_quay
	INNER JOIN data_da_rada_a_ormeggio
	ON trips_logs.fk_control_unit_data = data_da_rada_a_ormeggio.fk_control_unit_data 	
	WHERE ts_main_event_field_val BETWEEN $2 AND $3
	AND trips_logs.data_table_id::INTEGER = data_da_rada_a_ormeggio.id_data_da_rada_a_ormeggio
	AND trips_logs.fk_portinformer = $1
	AND fk_maneuvering IS NOT NULL
	AND maneuverings.fk_stop_quay != 0
	ORDER BY ts_main_event_field_val ASC
) AS position_data
ON id_control_unit_data = position_data.id_trip_data
WHERE control_unit_data.fk_portinformer = $1 
AND ts_fine_ormeggio BETWEEN $2 AND $3
)

--name: roadstead-register
(SELECT id_control_unit_data, ship_description, type_description, ts_anchor_drop,
countries.name as country, width, length, gross_tonnage, net_tonnage, position_data.roadstead 
FROM control_unit_data
INNER JOIN data_arrivo_in_rada
ON fk_control_unit_data = id_control_unit_data
INNER JOIN ships
ON fk_ship = id_ship
INNER JOIN ship_types
ON fk_ship_type = id_ship_type
INNER JOIN countries
ON fk_country_flag = id_country
LEFT JOIN (
	SELECT trips_logs.fk_control_unit_data AS id_trip_data, anchorage_points.description AS roadstead
	FROM trips_logs 
	INNER JOIN maneuverings
	ON trips_logs.fk_maneuvering = id_maneuvering
	INNER JOIN anchorage_points 
	ON maneuverings.fk_stop_anchorage_point = id_anchorage_point
	INNER JOIN data_arrivo_in_rada
	ON trips_logs.fk_control_unit_data = data_arrivo_in_rada.fk_control_unit_data 	
	WHERE ts_main_event_field_val BETWEEN $2 AND $3
	AND trips_logs.data_table_id::INTEGER = data_arrivo_in_rada.id_data_arrivo_in_rada
	AND trips_logs.fk_portinformer = $1
	AND fk_maneuvering IS NOT NULL
	AND maneuverings.fk_stop_anchorage_point != 0
	ORDER BY ts_main_event_field_val ASC
) AS position_data
ON id_control_unit_data = position_data.id_trip_data
WHERE control_unit_data.fk_portinformer = $1 
AND ts_anchor_drop BETWEEN $2 AND $3)
UNION
(SELECT id_control_unit_data, 
	ship_description, type_description, ts_anchor_drop,
	countries.name as country, width, length, gross_tonnage, net_tonnage, position_data.roadstead 
	FROM control_unit_data
	INNER JOIN data_da_ormeggio_a_rada
	ON fk_control_unit_data = id_control_unit_data
	INNER JOIN ships
	ON fk_ship = id_ship
	INNER JOIN ship_types
	ON fk_ship_type = id_ship_type
	INNER JOIN countries
	ON fk_country_flag = id_country
	LEFT JOIN (
		SELECT trips_logs.fk_control_unit_data AS id_trip_data, anchorage_points.description AS roadstead
		FROM trips_logs 
		INNER JOIN maneuverings
		ON trips_logs.fk_maneuvering = id_maneuvering
		INNER JOIN anchorage_points 
		ON maneuverings.fk_stop_anchorage_point = id_anchorage_point
		INNER JOIN data_da_ormeggio_a_rada
		ON trips_logs.fk_control_unit_data = data_da_ormeggio_a_rada.fk_control_unit_data 	
		WHERE ts_main_event_field_val BETWEEN $2 AND $3
		AND trips_logs.data_table_id::INTEGER = data_da_ormeggio_a_rada.id_data_da_ormeggio_a_rada
		AND trips_logs.fk_portinformer = $1
		AND fk_maneuvering IS NOT NULL
		AND maneuverings.fk_stop_anchorage_point != 0
		ORDER BY ts_main_event_field_val ASC
) AS position_data
ON id_control_unit_data = position_data.id_trip_data
	WHERE control_unit_data.fk_portinformer = $1 
	AND ts_anchor_drop BETWEEN $2 AND $3)
UNION
(
SELECT id_control_unit_data, 
ship_description, type_description, ts_anchor_drop,
countries.name as country, width, length, gross_tonnage, net_tonnage, position_data.roadstead 
FROM control_unit_data
INNER JOIN data_da_rada_a_rada
ON fk_control_unit_data = id_control_unit_data
INNER JOIN ships
ON fk_ship = id_ship
INNER JOIN ship_types
ON fk_ship_type = id_ship_type
INNER JOIN countries
ON fk_country_flag = id_country
	LEFT JOIN (
		SELECT trips_logs.fk_control_unit_data AS id_trip_data, anchorage_points.description AS roadstead
		FROM trips_logs 
		INNER JOIN maneuverings
		ON trips_logs.fk_maneuvering = id_maneuvering
		INNER JOIN anchorage_points 
		ON maneuverings.fk_stop_anchorage_point = id_anchorage_point
		INNER JOIN data_da_rada_a_rada
		ON trips_logs.fk_control_unit_data = data_da_rada_a_rada.fk_control_unit_data 	
		WHERE ts_main_event_field_val BETWEEN $2 AND $3
		AND trips_logs.data_table_id::INTEGER = data_da_rada_a_rada.id_data_da_rada_a_rada
		AND trips_logs.fk_portinformer = $1
		AND fk_maneuvering IS NOT NULL
		AND maneuverings.fk_stop_anchorage_point != 0
		ORDER BY ts_main_event_field_val ASC
) AS position_data
ON id_control_unit_data = position_data.id_trip_data
WHERE control_unit_data.fk_portinformer = $1 
AND ts_anchor_drop BETWEEN $2 AND $3
)

--name: departures-register
SELECT id_control_unit_data AS id_trip, 
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
AND ts_out_of_sight::TIMESTAMP BETWEEN $3 AND $4


--name: shipped-goods-register
SELECT id_control_unit_data AS id_trip,
ships.ship_description AS ship_name,
CASE WHEN quantity = '' THEN '0' ELSE quantity END,
unit, goods_categories.description AS goods_category,
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
INNER JOIN data_avvistamento_nave
ON data_avvistamento_nave.fk_control_unit_data = id_control_unit_data
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
WHERE control_unit_data.fk_portinformer = $1
AND ts_avvistamento BETWEEN $2 AND $3;

--name: traffic-list-register
SELECT
control_unit_data.id_control_unit_data AS id_trip,
ships.ship_description AS ship_name, 
ts_avvistamento AS ts_sighting,
num_container, num_passengers, num_camion, 
num_furgoni, num_rimorchi, num_auto, num_moto, num_camper, tons,
num_bus, num_minibus, traffic_list_mvnt_type, traffic_list_categories.description,
quays.description AS quay
FROM traffic_list INNER JOIN control_unit_data
ON traffic_list.fk_control_unit_data = id_control_unit_data
INNER JOIN traffic_list_categories
ON fk_traffic_list_category = id_traffic_list_category
INNER JOIN data_avvistamento_nave
ON data_avvistamento_nave.fk_control_unit_data = id_control_unit_data
INNER JOIN ships
ON control_unit_data.fk_ship = id_ship
INNER JOIN maneuverings
ON maneuverings.fk_control_unit_data = control_unit_data.id_control_unit_data
AND maneuverings.fk_state = 17
INNER JOIN quays
ON maneuverings.fk_stop_quay = quays.id_quay
WHERE control_unit_data.fk_portinformer = $1
AND ts_avvistamento BETWEEN $2 AND $3;