package spire

import (
	"fmt"
	"github.com/Akkadius/spire/internal/crashreporting"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/models"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/k0kubun/pp/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

type CrashAnalyticsFingerprintBackfillCommand struct {
	logger   *logrus.Logger
	command  *cobra.Command
	db       *database.DatabaseResolver
	pathmgmt *pathmgmt.PathManagement
}

func (c *CrashAnalyticsFingerprintBackfillCommand) Command() *cobra.Command {
	return c.command
}

func NewCrashAnalyticsCommand(
	logger *logrus.Logger,
	pathmgmt *pathmgmt.PathManagement,
	db *database.DatabaseResolver,
) *CrashAnalyticsFingerprintBackfillCommand {
	i := &CrashAnalyticsFingerprintBackfillCommand{
		logger:   logger,
		pathmgmt: pathmgmt,
		db:       db,
		command: &cobra.Command{
			Use:   "spire:crash-fp-backfill",
			Short: "Run the crash fingerprint backfill process (this is a one time process)",
			// unique to running within hosted
			Hidden: len(os.Getenv("GITHUB_CLIENT_ID")) == 0,
		},
	}

	// flag for ignoring existing fingerprints
	i.command.Flags().Bool("ignore-existing", false, "Ignore existing fingerprints and re-calculate them")

	i.command.Run = i.Handle

	return i
}

func (c *CrashAnalyticsFingerprintBackfillCommand) Handle(_ *cobra.Command, args []string) {
	fmt.Println("Running crash fingerprint backfill process...")

	// query the database for all crashes
	// for each crash, generate a fingerprint
	// update the crash with the fingerprint
	// this will allow us to group crashes by fingerprint
	// and then we can use the fingerprint to determine
	// if the crash is a duplicate or not

	var crashes []models.CrashReport
	c.db.GetSpireDb().Model(&models.CrashReport{}).Find(&crashes)

	// get ignore existing flag
	ignoreExisting, _ := c.command.Flags().GetBool("ignore-existing")

	for _, crash := range crashes {
		if crash.Fingerprint != "" && !ignoreExisting {
			continue
		}

		pp.Println("Reading crash report", crash.ID)
		fingerprint := crashreporting.FingerPrint(crash.CrashReport)

		pp.Println("Fingerprint", fingerprint)

		c.db.GetSpireDb().
			Model(&models.CrashReport{}).
			Where("id = ?", crash.ID).
			Update("fingerprint", fingerprint)
	}
}
