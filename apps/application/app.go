package application

import (
	"context"
	"controle-grupo-danca/apps/config"
	"controle-grupo-danca/pkg/database/connections"

	"github.com/uptrace/bun"
)

type ctxKey struct{}


type App struct {
	connections *connections.Manager
	Config      *config.Backend
}

func New(cfg config.Backend) (App, error) {
	return App{
		connections: connections.New(cfg.Database),
		Config:      &cfg,
	}, nil
}

func Ctx(ctx context.Context) App {
	cf, _ := ctx.Value(ctxKey{}).(App)

	return cf
}

func (a App) DB() *bun.DB {
	return a.connections.Get("grito-sepe")
}

func (c App) WithContext(ctx context.Context) context.Context {
	if cf, ok := ctx.Value(ctxKey{}).(App); ok {
		if cf == c {
			return ctx
		}
	}

	return context.WithValue(ctx, ctxKey{}, c)
}
