package controllers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type QueryController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewQueryController(db *database.DatabaseResolver, logger *logrus.Logger) *QueryController {
	return &QueryController{
		db:     db,
		logger: logger,
	}
}

func (q *QueryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "query/free-id-ranges/:table/:id", q.freeIdRanges, nil),
	}
}

// keep this code around for now because it handles dynamic fetching
func (q *QueryController) freeIdRangesOld(c echo.Context) error {
	db, err := q.db.GetEqemuDb().DB()
	if err != nil {
		q.logger.Warn(err)
	}

	sql := fmt.Sprintf(`
		SELECT (t1.id + 1) as gap_starts_at, 
			   (SELECT MIN(t3.id) -1 FROM %v t3 WHERE t3.id > t1.id) as gap_ends_at
		FROM %v t1
		WHERE NOT EXISTS (SELECT t2.id FROM %v t2 WHERE t2.id = t1.id + 1)
		HAVING gap_ends_at IS NOT NULL;
		`,
		c.Param("table"),
		c.Param("table"),
		c.Param("table"),
	)

	rows, err := db.Query(sql)
	if err != nil {
		q.logger.Warn(err)
	}
	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	resultRows := []map[string]string{}
	for rows.Next() {
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)

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

	return c.JSON(http.StatusOK, echo.Map{"data": resultRows})
}

type StartEndRange struct {
	StartId string `json:"start_id"`
	EndId   string `json:"emd_id"`
}

// this version is far faster than its query counterpart
func (q *QueryController) freeIdRanges(c echo.Context) error {

	// database instance
	db, err := q.db.GetEqemuDb().DB()
	if err != nil {
		q.logger.Warn(err)
	}

	// params
	table := c.Param("table")
	IdColumn := c.Param("id")

	// query
	query := fmt.Sprintf(`SELECT %v FROM %v ORDER BY %v`,
		IdColumn,
		table,
		IdColumn,
	)

	rows, err := db.Query(query)
	if err != nil {
		q.logger.Warn(err)
	}

	defer rows.Close()

	if err != nil {
		q.logger.Warn(err)
	}

	// scan ids
	ids := []string{}
	idMap := map[string]bool{}
	for rows.Next() {
		var Id string
		err := rows.Scan(&Id)
		if err != nil {
			q.logger.Fatal(err)
		}

		ids = append(ids, Id)
		idMap[Id] = true
	}

	if len(ids) > 0 {
		lowestId, err := strconv.ParseInt(ids[0], 10, 64)
		if err != nil {
			q.logger.Warn(err)
		}

		highestId, err := strconv.ParseInt(ids[len(ids)-1], 10, 64)
		if err != nil {
			q.logger.Warn(err)
		}

		startId := int64(0)
		endId := int64(0)
		ranges := []StartEndRange{}
		for i := lowestId; i <= highestId; i++ {

			// if id doesn't exist
			if _, ok := idMap[fmt.Sprintf("%v", i)]; !ok {
				if startId == 0 {
					startId = i
				}
			}

			// if id does exist
			if _, ok := idMap[fmt.Sprintf("%v", i)]; ok {
				if startId > 0 {
					// set last id
					endId = i - 1

					// append
					ranges = append(ranges, StartEndRange{
						StartId: fmt.Sprintf("%v", startId),
						EndId:   fmt.Sprintf("%v", endId),
					})

					// reset
					startId = 0
					endId = 0
				}
			}
		}

		return c.JSON(http.StatusOK, echo.Map{"data": ranges})
	}

	return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Ranges not found"})
}
