package database

import (
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/models"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// Connections application database connections
type Connections struct {
	spireDb *gorm.DB
	eqemuDb *gorm.DB
	logger  *logger.AppLogger
}

func (c *Connections) EqemuDb() *gorm.DB {
	return c.eqemuDb
}

func (c *Connections) SpireDb() *gorm.DB {
	return c.spireDb
}

func (c *Connections) SpireDbNoLog() *gorm.DB {
	return c.spireDb.Session(&gorm.Session{Logger: gormLogger.Default.LogMode(gormLogger.Silent)})
}

func NewConnections(
	spire *gorm.DB,
	EQEmu *gorm.DB,
	logger *logger.AppLogger,
) *Connections {
	return &Connections{
		spireDb: spire,
		eqemuDb: EQEmu,
		logger:  logger,
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
	&models.CrashReport{},
	&models.Setting{},
}

func (c *Connections) SpireMigrate(drop bool) error {
	for _, table := range spireTables {
		if drop {
			c.logger.Info().Msgf("Dropping table [%v]\n", table.TableName())
			_ = c.SpireDb().Migrator().DropTable(table)
		}

		// build migrator instance
		migrator := c.SpireDb().
			Session(&gorm.Session{Logger: gormLogger.Default.LogMode(gormLogger.Silent)}).
			Migrator()

		// only emit creation message when the table doesn't actually exist
		if !migrator.HasTable(table) {
			c.logger.Info().Msgf("[Database] Creating table [%v]", table.TableName())
		} else {
			c.logger.DebugVvv().Msgf("[Database] Already has table [%v]", table.TableName())
		}

		// always run migration incase there are schema changes
		err := migrator.AutoMigrate(table)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Connections) GetMigrationTables() []string {
	var tables []string
	for _, table := range spireTables {
		tables = append(tables, table.TableName())
	}

	return tables
}
