package clientfiles

import (
	"database/sql"
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
)

type Exporter struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewExporter(db *gorm.DB, logger *logrus.Logger) *Exporter {
	return &Exporter{db: db, logger: logger}
}

func (e *Exporter) getDatabase() *sql.DB {
	// get database instance
	db, err := e.db.DB()
	if err != nil {
		e.logger.Fatal(err)
	}

	return db
}

func (e *Exporter) ExportSpells() string {
	var entries []map[string]interface{}
	e.db.Model(&models.SpellsNew{}).Find(&entries)

	columns, err := database.GetTableSchema(e.getDatabase(), "spells_new")
	if err != nil {
		e.logger.Fatal(err)
	}

	var rows []string
	for _, e := range entries {
		var cols []string
		for _, column := range columns {
			cols = append(cols, fmt.Sprintf("%v", e[column.Column]))
		}

		rows = append(rows, strings.Join(cols, "^"))
	}

	return strings.Join(rows, "\n")
}

func (e *Exporter) ExportDbStr() string {
	var entries []map[string]interface{}
	e.db.Model(&models.DbStr{}).Find(&entries)

	columns, err := database.GetTableSchema(e.getDatabase(), "db_str")
	if err != nil {
		e.logger.Fatal(err)
	}

	var rows []string
	for _, e := range entries {
		var cols []string
		for _, column := range columns {
			cols = append(cols, fmt.Sprintf("%v", e[column.Column]))
		}
		cols = append(cols, "0")

		rows = append(rows, strings.Join(cols, "^"))
	}

	return strings.Join(rows, "\n")
}
