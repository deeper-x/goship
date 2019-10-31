package ldb

import (
	"database/sql"
	"log"

	"github.com/deeper-x/goship/conf"
)

// GetActiveStations todo doc
func GetActiveStations() []map[string]string {
	var portinformerCode, idPortinformer, tsFirstCreated, isActive sql.NullString
	var result = []map[string]string{}

	connector := Connect()

	mapper.GenResource(conf.PMeteoSQL)
	rows, err := mapper.resource.Query(connector, "meteo-active-stations")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&idPortinformer,
			&portinformerCode,
			&tsFirstCreated,
			&isActive,
		)

		if err != nil {
			log.Fatal(err)
		}

		tmpDict := map[string]string{
			"id_portinformer":   idPortinformer.String,
			"portinformer_code": portinformerCode.String,
			"ts_first_created":  tsFirstCreated.String,
			"is_active":         isActive.String,
		}

		result = append(result, tmpDict)
	}

	return result

}
