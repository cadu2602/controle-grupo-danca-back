package main

import (
	"controle-grupo-danca/apps/application"
	"controle-grupo-danca/apps/cmd/migrate"
	"controle-grupo-danca/apps/config"
	"os"
	"sort"

	"github.com/rs/zerolog/log"
	zero "github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Description: "Grupo dança",
		Usage:       "Grupo dança Backend",
		Commands: []*cli.Command{
			migrate.MakeCmd,
			migrate.MigrateCmd,
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "load configuration from",
				DefaultText: "",
			},
			&cli.StringFlag{
				Name:        "level",
				Aliases:     []string{"l"},
				Usage:       "define log level",
				DefaultText: "info",
			},
		},
		Before: func(cmd *cli.Context) error {
			cfgPath := cmd.String("config")
			if cfgPath == "" {
				cfgPath = "backend.yml"
			}

			cfg, err := config.Load(cfgPath)
			if err != nil {
				return err
			}



			cmd.Context = cfg.WithContext(cmd.Context)

			log.Debug().Msg("config loaded")

			app, err := application.New(cfg)

			if err != nil {
				return err
			}

			cmd.Context = app.WithContext(cfg.WithContext(cmd.Context))

			return nil
		},
	}


	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)

	if err != nil {
		zero.Fatal().Err(err).Msg("fail run application")
	}
}
