package controllers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
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
		routes.RegisterRoute(http.MethodGet, "query/schema/table/:table", q.getTableSchema, nil),
		routes.RegisterRoute(http.MethodGet, "query/free-id-ranges/:table/:id", q.freeIdRanges, nil),
		routes.RegisterRoute(http.MethodGet, "query/free-ids-reserved/:table/:id/:name", q.freeIdsReserved, nil),
		routes.RegisterRoute(http.MethodGet, "query/expansion-stats", q.expansionStats, nil),
	}
}

type StartEndRange struct {
	StartId string `json:"start_id"`
	EndId   string `json:"end_id"`
}

// this version is far faster than its query counterpart
// freeIdRanges gets free contiguous blocks of ids
func (q *QueryController) freeIdRanges(c echo.Context) error {

	// params
	table := c.Param("table")
	IdColumn := c.Param("id")

	// database instance
	db, err := q.db.Get(q.getModelFromString(table), c).DB()
	if err != nil {
		q.logger.Warn(err)
	}

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

	maxId := 0

	// scan ids
	ids := []string{}
	idMap := map[string]bool{}
	for rows.Next() {
		var Id string
		err := rows.Scan(&Id)
		if err != nil {
			q.logger.Warn(err)
		}

		ids = append(ids, Id)
		idMap[Id] = true

		// set max id
		id, _ := strconv.Atoi(Id)
		maxId = id
	}

	gapSize := 10

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
					if (endId - startId) > int64(gapSize) {
						ranges = append(ranges, StartEndRange{
							StartId: fmt.Sprintf("%v", startId),
							EndId:   fmt.Sprintf("%v", endId),
						})
					}

					// reset
					startId = 0
					endId = 0
				}
			}
		}

		// if there are no free ranges, use the max ID
		if len(ranges) == 0 {
			ranges = append(ranges, StartEndRange{
				StartId: fmt.Sprintf("%v", maxId+1),
				EndId:   fmt.Sprintf("%v", maxId+1),
			})
		}

		return c.JSON(http.StatusOK, echo.Map{"data": ranges})
	}

	return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Ranges not found"})
}

func (q *QueryController) getTableSchema(c echo.Context) error {
	table := c.Param("table")
	db := q.db.Get(q.getModelFromString(table), c)
	schema, err := database.GetTableSchema(db, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"data": schema})
}

func (q *QueryController) freeIdsReserved(c echo.Context) error {
	table := c.Param("table")

	db, err := q.db.Get(q.getModelFromString(table), c).DB()
	if err != nil {
		q.logger.Warn(err)
	}

	IdColumn := c.Param("id")
	keyName := c.Param("name")

	query := fmt.Sprintf(`select %v, %v from %v where %v like '%%placeholder%%' or %v like '%% reserved%%' order by %v`,
		IdColumn,
		keyName,
		table,
		keyName,
		keyName,
		IdColumn,
	)

	return c.JSON(http.StatusOK, echo.Map{"data": database.GenericQuery(db, query)})
}

func (q *QueryController) getModelFromString(s string) models.Modelable {
	if s == "zone" {
		return models.Zone{}
	}
	if s == "items" {
		return models.Item{}
	}
	if s == "tasks" {
		return models.Task{}
	}
	if s == "spells_new" {
		return models.SpellsNew{}
	}

	return models.Zone{}
}

func (q *QueryController) expansionStats(c echo.Context) error {
	db, err := q.db.Get(q.getModelFromString("zone"), c).DB()
	if err != nil {
		q.logger.Warn(err)
	}

	// gather content tables
	query := fmt.Sprintf(`SELECT TABLE_NAME, COLUMN_NAME, DATA_TYPE, IS_NULLABLE, COLUMN_DEFAULT FROM INFORMATION_SCHEMA.COLUMNS WHERE column_name LIKE 'min_expansion'`)
	tableNames := []string{}
	for _, m := range database.GenericQuery(db, query) {
		tableNames = append(tableNames, m["TABLE_NAME"])
	}

	response := map[string][]map[string]string{}

	for _, name := range tableNames {
		q := fmt.Sprintf("select count(*) as count, min_expansion FROM %v WHERE `min_expansion` > -1 GROUP BY `min_expansion`", name)
		response[name] = database.GenericQuery(db, q)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": response})
}
