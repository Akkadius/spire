package database

import (
	"fmt"
	"github.com/Akkadius/spire/models"
	"github.com/jinzhu/gorm"
)

// Connections application database connections
type Connections struct {
	spireDb *gorm.DB
	eqemuDb *gorm.DB
}

func (c Connections) EqemuDb() *gorm.DB {
	return c.eqemuDb
}

func (c Connections) SpireDb() *gorm.DB {
	return c.spireDb
}

func NewConnections(
	spire *gorm.DB,
	EQEmu *gorm.DB) *Connections {
	return &Connections{
		spireDb: spire,
		eqemuDb: EQEmu,
	}
}

var spireTables = []models.Modelable{
	&models.User{},
	&models.UserServerDatabaseConnection{},
	&models.ServerDatabaseConnection{},
	&models.AnalyticEvent{},
	&models.AnalyticEventCount{},
}

func (c Connections) SpireMigrate(drop bool) {
	for _, table := range spireTables {
		if drop {
			fmt.Printf("Dropping table [%v]\n", table.TableName())
			c.SpireDb().DropTable(table)
		}
		fmt.Printf("Migrating table [%v]\n", table.TableName())
		c.SpireDb().AutoMigrate(table)

		indexes, ok := table.(models.Indexable)
		if ok {
			fmt.Printf("Running indexes for [%v]\n", table.TableName())

			for indexName, indexKeys := range indexes.Indexes() {
				c.SpireDb().Model(table).AddIndex(indexName, indexKeys...)
				fmt.Printf("Adding index for [%v] index [%v] keys %v\n", table.TableName(), indexName, indexKeys)
			}
		}
	}
}
