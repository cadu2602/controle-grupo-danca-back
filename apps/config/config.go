package config

import (
	"context"
	"controle-grupo-danca/pkg/database"
	"reflect"
)

type ctxKey struct{}

type Backend struct {
	Port       int                 `fig:"port" yaml:"port"`
	Database   database.Config     `fig:"database" yaml:"database"`
}

func Ctx(ctx context.Context) Backend {
	cf, _ := ctx.Value(ctxKey{}).(Backend)

	return cf
}


func (c Backend) WithContext(ctx context.Context) context.Context {
	if cf, ok := ctx.Value(ctxKey{}).(Backend); ok {
		if reflect.DeepEqual(cf, c) {
			return ctx
		}
	}

	return context.WithValue(ctx, ctxKey{}, c)
}
