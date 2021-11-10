package main

import (
	"fmt"
	configs "project/mock_api/config"
	"project/mock_api/controllers"
	"project/mock_api/middlewares"
	"project/mock_api/models"
	"project/mock_api/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	config := configs.GetConfig()

	db := configs.MysqlConnect(config)

	userModel := models.NewUserModel(db)
	bookModel := models.NewBookModel(db)

	userController := controllers.NewController(userModel)
	bookController := controllers.NewBookController(bookModel)
	// controllers.SetUp()
	// controllers.TestReq()

	ec := echo.New()

	routes.RegisterPath(ec, userController, bookController)
	middlewares.GlobalMiddlewares(ec)

	runServer := fmt.Sprint("localhost: ", config.Port)

	if err := ec.Start(runServer); err != nil {
		log.Info("shutting down the server")
	}
}
