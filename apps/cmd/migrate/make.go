package migrate

import (
	"controle-grupo-danca/migrations"
	"strings"

	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"
)

func getMigrator(conn *bun.DB) *migrate.Migrator {
	return migrate.NewMigrator(conn, migrations.Migrations)
}

var MakeCmd = &cli.Command{
	Name:  "migrate:make",
	Usage: "create up and down migrations",
	Action: func(cmd *cli.Context) error {
		name := strings.Join(cmd.Args().Slice(), "_")
		logger := zerolog.Ctx(cmd.Context)

		migrator := getMigrator(nil)

		isSQL := cmd.Bool("sql")

		if isSQL {
			logger.Info().Msg("creating sql migration...")

			files, err := migrator.CreateSQLMigrations(cmd.Context, name)

			if err != nil {
				return err
			}

			for _, mf := range files {
				logger.Info().Msgf("created migration %s (%s)\n", mf.Name, mf.Path)
			}

			return nil
		}

		file, err := migrator.CreateGoMigration(cmd.Context, name)

		if err != nil {
			return err
		}

		logger.Info().Msgf("created migration %s (%s)\n", file.Name, file.Path)

		return nil
	},
}
