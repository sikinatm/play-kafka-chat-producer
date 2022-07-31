//go:generate go run github.com/google/wire/cmd/wire
//go:build wireinject
// +build wireinject

package main

import (
	"github/Shitomo/producer/adapter/controller"
	"github/Shitomo/producer/adapter/gateway"
	"github/Shitomo/producer/adapter/presenter"
	"github/Shitomo/producer/adapter/producer"
	"github/Shitomo/producer/driver/server"
	"github/Shitomo/producer/usecase/interactor"

	"github.com/google/wire"
)

var producerSet = wire.NewSet(
	producer.NewUserProducer,
)

var gatewaySet = wire.NewSet(
	gateway.NewUserGateway,
)

var presenterSet = wire.NewSet(
	presenter.NewUserRegisterPresenter,
)

var interactorSet = wire.NewSet(
	interactor.NewUserRegisterInteractor,
)

var controllerSet = wire.NewSet(
	controller.NewUserRegisterController,
)

func InitializeHTTPServer() *server.HTTPServer {
	wire.Build(
		producerSet,
		gatewaySet,
		presenterSet,
		interactorSet,
		controllerSet,
		server.NewHTTPServer,
	)

	return nil
}
