--name: all-active-trips
SELECT DISTINCT ON (id_control_unit_data) id_control_unit_data,
ships.ship_description AS ship_name,
ship_types.type_description||'('||ship_types.type_acronym||')' AS ship_type,
ships.length AS length, ships.width AS width, ships.gross_tonnage AS gross_tonnage,
ships.net_tonnage AS net_tonnage,
ship_current_activities.description AS current_activity,
last_trip_ts AS timestamp,
shipped_goods_data.shipped_goods_row AS shipping,
agency_data.agency
FROM control_unit_data INNER JOIN ships
ON fk_ship = id_ship
INNER JOIN ( 
	SELECT DISTINCT ON (fk_control_unit_data) agencies.description AS agency, fk_control_unit_data
	FROM trips_logs
	INNER JOIN agencies
	ON trips_logs.fk_agency = id_agency
	ORDER BY fk_control_unit_data, ts_main_event_field_val DESC
) AS agency_data
ON agency_data.fk_control_unit_data = id_control_unit_data
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
INNER JOIN ship_current_activities
ON fk_ship_current_activity = id_activity
INNER JOIN (
    SELECT fk_control_unit_data, MAX(ts_main_event_field_val) AS last_trip_ts
        FROM trips_logs INNER JOIN control_unit_data
        ON id_control_unit_data = fk_control_unit_data
        WHERE control_unit_data.fk_portinformer = $1
	AND fk_state = ANY('{13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 25, 26, 27, 28}')
        GROUP BY fk_control_unit_data
    ) as last_trip_log
 ON last_trip_log.fk_control_unit_data = id_control_unit_data
 INNER JOIN ship_types
 ON ships.fk_ship_type = ship_types.id_ship_type
 WHERE is_active  = true
 AND fk_portinformer = $1
 ORDER BY id_control_unit_data, timestamp DESC;


--name: all-anchored
SELECT 
DISTINCT ON (id_control_unit_data) id_control_unit_data, 
ship_description,
type_acronym as ship_type, 
RES.ts_anchoring_time, 
ship_current_activities.description AS current_activity,
anchorage_points.description AS anchorage_point,
iso3, 
gross_tonnage, 
ships.length, 
ships.width,
shipped_goods_data.shipped_goods_row AS shipped_goods_data,
data_previsione_arrivo_nave.ts_mooring_time AS ts_planned_mooring,
agencies.description AS agency
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
	SELECT fk_control_unit_data, MAX(ts_main_event_field_val) AS ts_anchoring_time
	FROM trips_logs
	WHERE fk_portinformer = $1
	AND fk_state IN (16, 19, 27)
	GROUP BY fk_control_unit_data, fk_portinformer
	) AS RES
ON id_control_unit_data = RES.fk_control_unit_data
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
LEFT JOIN data_previsione_arrivo_nave
ON data_previsione_arrivo_nave.fk_control_unit_data = control_unit_data.id_control_unit_data
INNER JOIN agencies
ON agencies.id_agency = data_previsione_arrivo_nave.fk_agency
WHERE fk_ship_current_activity = 2
AND is_active = true 
AND control_unit_data.fk_portinformer = $2
ORDER BY id_control_unit_data, RES.ts_anchoring_time DESC;

--name: all-moored
SELECT DISTINCT ON (id_control_unit_data) id_control_unit_data, ship_description, 
RES.ts_fine_ormeggio, 
ship_current_activities.description AS current_activity, 
quays.description AS quay,
berths.description AS berth,
shipped_goods_data.shipped_goods_row AS shipped_goods_data,
iso3, gross_tonnage, ships.length, ships.width, type_acronym,
agency_data.agency AS agency
FROM control_unit_data 
INNER JOIN ships
ON fk_ship = id_ship
INNER JOIN ship_current_activities
ON fk_ship_current_activity = id_activity
INNER JOIN latest_maneuverings
ON latest_maneuverings.fk_control_unit_data = id_control_unit_data
INNER JOIN quays
ON latest_maneuverings.fk_stop_quay = id_quay
INNER JOIN berths
ON latest_maneuverings.fk_stop_berth = id_berth
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
SELECT fk_control_unit_data, MAX(ts_main_event_field_val) AS ts_fine_ormeggio
FROM trips_logs
WHERE fk_portinformer = $1
AND fk_state in (17, 18, 20, 21, 22)
GROUP BY fk_control_unit_data, fk_portinformer
) AS RES
ON id_control_unit_data = RES.fk_control_unit_data
INNER JOIN ( 
	SELECT DISTINCT ON (fk_control_unit_data) agencies.description AS agency, fk_control_unit_data
	FROM trips_logs
	INNER JOIN agencies
	ON trips_logs.fk_agency = id_agency
	ORDER BY fk_control_unit_data, ts_main_event_field_val DESC
) AS agency_data
ON agency_data.fk_control_unit_data = id_control_unit_data
INNER JOIN countries
ON countries.id_country = ships.fk_country_flag
INNER JOIN ship_types
ON ships.fk_ship_type = ship_types.id_ship_type
WHERE fk_ship_current_activity = 5
AND control_unit_data.is_active = true 
AND control_unit_data.fk_portinformer = $2
ORDER BY id_control_unit_data, RES.ts_fine_ormeggio DESC;


--name: arrival-previsions
SELECT
ship_description AS ship, 
ts_arrival_prevision,
ship_types.type_acronym AS ship_type,  
countries.iso3 AS ship_flag,
ships.width AS ship_width,
ships.length AS ship_length,
ships.gross_tonnage AS gross_tonnage,
ships.net_tonnage AS net_tonnage,
draft_aft, 
draft_fwd,
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
AND planned_arrivals.fk_portinformer = $1
ORDER BY ts_arrival_prevision DESC;


--name: shiftings
SELECT 
DISTINCT ON (id_control_unit_data) id_control_unit_data AS id_trip, 
ts_main_event_field_val AS ts_shifting, 
imo, 
ship_description AS ship_name, 
type_acronym AS ship_type, 
iso3 AS country, 
start_quay.description||'/'||start_berth.description as FROM_QUAY,
stop_quay.description||'/'||stop_berth.description as TO_QUAY,
start_anchorage.description as FROM_ANCH,
stop_anchorage.description as TO_ANCH,
agency_data.agency AS agency
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
INNER JOIN ( 
	SELECT DISTINCT ON (fk_control_unit_data) agencies.description AS agency, fk_control_unit_data
	FROM trips_logs
	INNER JOIN agencies
	ON trips_logs.fk_agency = id_agency
	ORDER BY fk_control_unit_data, ts_main_event_field_val DESC
) AS agency_data
ON agency_data.fk_control_unit_data = id_control_unit_data
WHERE control_unit_data.fk_portinformer = $1
AND trips_logs.fk_state IN (18, 19, 20, 21, 22, 27)
AND LENGTH(ts_main_event_field_val) = 16
AND ts_main_event_field_val::date = now()::date
ORDER BY id_control_unit_data, ts_main_event_field_val DESC;



--name: shifting-previsions
SELECT ship_description AS ship, 
ts_shifting_prevision,
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
AND planned_shiftings.fk_portinformer = $1
ORDER BY ts_shifting_prevision DESC;


--name: departure-previsions
SELECT ship_description AS ship, 
ts_departure_prevision,
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
AND planned_departures.fk_portinformer = $1
ORDER BY ts_departure_prevision DESC;

--name: arrivals
SELECT DISTINCT ON (id_control_unit_data) id_control_unit_data AS id_trip, 
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
quays.description||'-'||berths.description AS destination_quay_berth,
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
AND ts_avvistamento::DATE = current_date
ORDER BY id_control_unit_data;

--name: departures
SELECT DISTINCT ON (id_control_unit_data) id_control_unit_data AS id_trip, 
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
AND ts_out_of_sight::DATE = current_date
ORDER BY id_control_unit_data, ts_out_of_sight DESC;


--name: shipped-goods
SELECT
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
macro_categories.description AS macro_category,
goods_mvmnt_type
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
WHERE control_unit_data.fk_portinformer = $1
AND control_unit_data.is_active = true;


--name: traffic-list
SELECT control_unit_data.id_control_unit_data AS id_trip,
ships.ship_description AS ship_name, 
num_container, 
num_passengers, 
num_camion, 
num_furgoni, 
num_rimorchi, 
num_auto, 
num_moto, 
num_camper, 
tons,
num_bus, 
num_minibus, 
traffic_list_mvnt_type, 
traffic_list_categories.description,
quays.description AS quay
FROM traffic_list LEFT JOIN control_unit_data
ON fk_control_unit_data = id_control_unit_data
LEFT JOIN traffic_list_categories
ON fk_traffic_list_category = id_traffic_list_category
RIGHT JOIN ships
ON control_unit_data.fk_ship = id_ship
LEFT JOIN maneuverings
ON maneuverings.fk_control_unit_data = control_unit_data.id_control_unit_data
AND maneuverings.fk_state = 17
LEFT JOIN quays
ON maneuverings.fk_stop_quay = quays.id_quay
WHERE control_unit_data.fk_portinformer = $1
AND control_unit_data.is_active = true;
