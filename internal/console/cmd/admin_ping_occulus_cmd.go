package cmd

import (
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/occulus"
	"github.com/k0kubun/pp/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type AdminPingOcculus struct {
	db           *gorm.DB
	logger       *logrus.Logger
	command      *cobra.Command
	serverconfig *eqemuserverconfig.Config
	occulus      *occulus.Proxy
}

func (c *AdminPingOcculus) Command() *cobra.Command {
	return c.command
}

func NewAdminPingOcculus(
	db *gorm.DB,
	logger *logrus.Logger,
	serverconfig *eqemuserverconfig.Config,
	occulus *occulus.Proxy,
) *AdminPingOcculus {
	i := &AdminPingOcculus{
		db:           db,
		logger:       logger,
		serverconfig: serverconfig,
		occulus:      occulus,
		command: &cobra.Command{
			Use:   "admin:ping-occulus",
			Short: "Pings Occulus",
		},
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *AdminPingOcculus) Handle(cmd *cobra.Command, args []string) {
	// login
	err := c.occulus.Login()
	if err != nil {
		c.logger.Error(err)
	}

	// test getting process counts
	counts, err := c.occulus.GetProcessCounts()
	if err != nil {
		c.logger.Error(err)
	}

	pp.Println(counts)
}

// Validate implementation of the Command interface
func (c *AdminPingOcculus) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
