//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"oss/internal/biz"
	"oss/internal/conf"
	"oss/internal/data"
	"oss/internal/pkg/minio"
	"oss/internal/pkg/task"
	"oss/internal/server"
	"oss/internal/service"
)

// wireApp init kratos application.
func wireApp(c *conf.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, task.ProviderSet, minio.ProviderSet, service.ProviderSet, newApp))
}
