--name: meteo-active-stations
SELECT ts_first_created, id_portinformer, portinformer_code,
CASE WHEN ts_first_created > NOW() - interval '1 hour' THEN 'true'
ELSE 'false'
END AS is_active
FROM live_meteo_data
INNER JOIN portinformers
ON fk_portinformer = id_portinformer;
