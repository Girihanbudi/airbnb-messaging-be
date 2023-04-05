package app

import (
	_ "airbnb-messaging-be/docs"

	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

func (a App) registerHttpHandler() {
	a.HttpServer.Router.GET("/docs/*any", ginswagger.WrapHandler(swaggerfiles.Handler))
}

func (a App) registerEventHandler() {
	a.SmsHandler.RegisterApi()
}
