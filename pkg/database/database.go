package database

import (
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"gopkg.in/mcuadros/go-defaults.v1"
)

type Config struct {
	Host     string `fig:"host" yaml:"host" default:"localhost"`
	Port     int    `fig:"port" yaml:"port" default:"5432"`
	Username string `fig:"username" yaml:"username" default:"mes"`
	Password string `fig:"password" yaml:"password"`
	DBName   string `fig:"dbname" yaml:"dbname" default:"mes"`
}

func ConfigFromEnv() (Config, error) {
	var cfg Config
	var err error

	defaults.SetDefaults(&cfg)

	return cfg, err
}

func (c Config) Connector() *pgdriver.Connector {

	return pgdriver.NewConnector(
		pgdriver.WithAddr(fmt.Sprintf("%s:%v", "localhost", "5440")),
		pgdriver.WithDatabase("grito-sepe"),
		pgdriver.WithUser("root"),
		pgdriver.WithPassword("root"),
		pgdriver.WithTLSConfig(nil),
	)
}

func NewDB(cfg Config) *bun.DB {
	sqldb := sql.OpenDB(cfg.Connector())

	db := bun.NewDB(sqldb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook())

	return db
}
