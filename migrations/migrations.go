package migrations

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/uptrace/bun/migrate"
)

// A collection of migrations.
var Migrations = migrate.NewMigrations()

func getLogger(ctx context.Context, name, direction string) zerolog.Logger {
	return zerolog.Ctx(ctx).With().
		Str("action", "migrate").
		Str("name", name).
		Str("direction", direction).
		Logger()
}
