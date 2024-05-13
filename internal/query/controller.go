package query

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Controller struct {
	db     *database.Resolver
	logger *logger.AppLogger
}

func NewController(db *database.Resolver, logger *logger.AppLogger) *Controller {
	return &Controller{
		db:     db,
		logger: logger,
	}
}

func (q *Controller) Routes() []*routes.Route {
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
func (q *Controller) freeIdRanges(c echo.Context) error {

	// params
	table := c.Param("table")
	IdColumn := c.Param("id")

	// database instance
	db, err := q.db.Get(q.getModelFromString(table), c).DB()
	if err != nil {
		q.logger.Warn().Err(err).Msg("Failed to get db")
	}

	// query
	query := fmt.Sprintf(`SELECT %v FROM %v ORDER BY %v`,
		IdColumn,
		table,
		IdColumn,
	)

	rows, err := db.Query(query)
	if err != nil {
		q.logger.Warn().Err(err).Msg("Failed to query")
	}

	defer rows.Close()

	if err != nil {
		q.logger.Warn().Err(err).Msg("Failed to query")
	}

	maxId := 0

	// scan ids
	var ids []string
	idMap := map[string]bool{}
	for rows.Next() {
		var Id string
		err := rows.Scan(&Id)
		if err != nil {
			q.logger.Warn().Err(err).Msg("Failed to scan")
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
			q.logger.Warn().Err(err).Msg("Failed to parse")
		}

		highestId, err := strconv.ParseInt(ids[len(ids)-1], 10, 64)
		if err != nil {
			q.logger.Warn().Err(err).Msg("Failed to parse")
		}

		startId := int64(0)
		endId := int64(0)
		var ranges []StartEndRange
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

func (q *Controller) getTableSchema(c echo.Context) error {
	table := c.Param("table")
	db := q.db.Get(q.getModelFromString(table), c)
	schema, err := database.GetTableSchema(db, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"data": schema})
}

func (q *Controller) freeIdsReserved(c echo.Context) error {
	table := c.Param("table")

	db, err := q.db.Get(q.getModelFromString(table), c).DB()
	if err != nil {
		q.logger.Warn().Err(err).Msg("Failed to get db")
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

	return c.JSON(http.StatusOK, echo.Map{"data": GenericQuery(db, query)})
}

func (q *Controller) getModelFromString(s string) models.Modelable {
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

func (q *Controller) expansionStats(c echo.Context) error {
	db, err := q.db.Get(q.getModelFromString("zone"), c).DB()
	if err != nil {
		q.logger.Warn().Err(err).Msg("Failed to get db")
	}

	// gather content tables
	query := fmt.Sprintf(`SELECT TABLE_NAME, COLUMN_NAME, DATA_TYPE, IS_NULLABLE, COLUMN_DEFAULT FROM INFORMATION_SCHEMA.COLUMNS WHERE column_name LIKE 'min_expansion'`)
	var tableNames []string
	for _, m := range GenericQuery(db, query) {
		tableNames = append(tableNames, m["TABLE_NAME"])
	}

	response := map[string][]map[string]string{}

	for _, name := range tableNames {
		q := fmt.Sprintf("select count(*) as count, min_expansion FROM %v WHERE `min_expansion` > -1 GROUP BY `min_expansion`", name)
		response[name] = GenericQuery(db, q)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": response})
}
