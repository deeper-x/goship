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

--name: moored-register
(SELECT id_control_unit_data, 
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
)

--name: roadstead-register
(SELECT id_control_unit_data, 
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