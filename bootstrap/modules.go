package bootstrap

import (
	"go.uber.org/fx"
	"tiny-url/api/controllers"
	"tiny-url/api/routes"
	"tiny-url/lib"
	"tiny-url/middlewares"
	"tiny-url/repository"
	"tiny-url/services"
)

var CommonModules = fx.Options(
	controllers.Module,
	routes.Module,
	lib.Module,
	services.Module,
	repository.Module,
	middlewares.Module,
)
