package config

import (
	"controle-grupo-danca/pkg/database"
	"controle-grupo-danca/pkg/errors"

	"path/filepath"

	"github.com/kkyr/fig"
)

func Load(file string) (cfg Backend, err error) {
	if file != "" {
		err = fig.Load(&cfg,
			fig.File(filepath.Base(file)),
			fig.Dirs(filepath.Dir(file)),
		)
		if err != nil {
			return cfg, errors.ErrFailToLoadConfig.WithErr(err)
		}

		return cfg, nil
	}

	cfg.Database, err = database.ConfigFromEnv()
	if err != nil {
		return cfg, err
	}

	return cfg, err
}
