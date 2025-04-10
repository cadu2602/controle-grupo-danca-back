package migrations

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

func init() {
	type Members struct {
		bun.BaseModel `bun:"table:members,alias:members"`
		ID            int64         `bun:"id,pk,autoincrement"`
		Name    			string        `bun:"name,notnull"`
		DateBirth     time.Time     `bun:"date_brith" json:"dateBirth,omitempty"`
		RG    			  string        `bun:"rg" json:"rg,omitempty"`
		CPF    			  string        `bun:"cpf" json:"cpf,omitempty"`
	}

	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		logger := getLogger(ctx, "add_table_members", "up")

		logger.Info().Msg("running...")

		_, err := db.NewCreateTable().
			Model(&Members{}).
			WithForeignKeys().
			Exec(ctx)
		if err != nil {
			return err
		}

		logger.Info().Msgf("done")
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		logger := getLogger(ctx, "machine operator", "down")

		logger.Warn().Msg("running...")

		_, err := db.NewDropTable().
			Model(&Members{}).Exec(ctx)

		logger.Info().Msgf("done")
		return err
	})
}
