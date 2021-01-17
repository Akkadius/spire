package database

import (
	"github.com/Akkadius/spire/models"
	"github.com/jinzhu/gorm"
)

// application database connections
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
}

func (c Connections) SpireMigrate(drop bool) {
	for _, table := range spireTables {
		if drop {
			c.SpireDb().DropTable(table)
		}
		c.SpireDb().AutoMigrate(table)
	}
}
