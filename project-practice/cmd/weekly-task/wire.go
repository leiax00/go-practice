//+build wireinject
//go:generate wire

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"project-practice/internal/weekly-task/biz"
	"project-practice/internal/weekly-task/conf"
	"project-practice/internal/weekly-task/data"
	"project-practice/internal/weekly-task/server"
	"project-practice/internal/weekly-task/service"
)

func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
