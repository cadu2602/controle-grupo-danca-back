package migrate

import (
	"controle-grupo-danca/apps/application"

	"github.com/rs/zerolog"
	"github.com/urfave/cli/v2"
)

var MigrateCmd = &cli.Command{
	Name:  "migrate:up",
	Usage: "Update database schema",
	Action: func(cmd *cli.Context) error {
		ctx := cmd.Context
		app := application.Ctx(ctx)

		logger := zerolog.Ctx(ctx)

		migrator := getMigrator(app.DB())

		err := migrator.Init(ctx)

		if err != nil {
			return err
		}

		group, err := migrator.Migrate(ctx)

		if err != nil {
			return err
		}

		if group.IsZero() {
			logger.Warn().Msg("there are no new migrations to run (database is up to date)")
			return nil
		}

		logger.Info().Msgf("migrated to %s", group)

		return nil
	},
}
