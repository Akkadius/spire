package query

import (
	"database/sql"
	"fmt"
	"log"
)

func GenericQuery(db *sql.DB, sql string) []map[string]string {
	rows, err := db.Query(sql)
	if err != nil {
		log.Println(err)
	}
	columns, err := rows.Columns()
	if err != nil {
		log.Println(err)
	}
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	resultRows := []map[string]string{}
	for rows.Next() {
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}
		err := rows.Scan(valuePtrs...)
		if err != nil {
			log.Println(err)
		}

		tmpStruct := map[string]string{}

		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			tmpStruct[col] = fmt.Sprintf("%s", v)
		}

		resultRows = append(resultRows, tmpStruct)
	}

	return resultRows
}
