package routes

import (
	"tiny-url/api/controllers"
	"tiny-url/lib"
	"tiny-url/middlewares"
)

type URLRoutes struct {
	logger        lib.Logger
	handler       lib.RequestHandler
	urlController controllers.URLController
	middleware    middlewares.Middlewares
}

func (u URLRoutes) Setup() {
	u.logger.Info("Setting up routes")
	api := u.handler.Gin.Group("/api")
	{
		api.GET("/url", u.urlController.GetUrl)
		api.GET("/url/:hash", u.urlController.GetOneUrl)
		api.PUT("/url", u.urlController.SaveUrl)
	}
}

// NewURLRoutes creates new user controller
func NewURLRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	userController controllers.URLController,
) URLRoutes {
	return URLRoutes{
		handler:       handler,
		logger:        logger,
		urlController: userController,
	}
}
