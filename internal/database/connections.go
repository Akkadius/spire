package database

import (
	"fmt"
	"github.com/Akkadius/spire/internal/models"
	"gorm.io/gorm"
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
	&models.UserEventLog{},
	&models.UserServerResourcePermission{},
}

func (c Connections) SpireMigrate(drop bool) {
	for _, table := range spireTables {
		if drop {
			fmt.Printf("Dropping table [%v]\n", table.TableName())
			_ = c.SpireDb().Migrator().DropTable(table)
		}
		fmt.Printf("Migrating table [%v]\n", table.TableName())
		_ = c.SpireDb().Migrator().AutoMigrate(table)
	}
}
