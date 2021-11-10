package routes

import (
	"project/mock_api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(ec *echo.Echo, uc *controllers.UserController, bc *controllers.BookController) {
	ec.GET("/users", uc.GetAllUser, middleware.JWT([]byte("rahasia")))
	ec.GET("/user", uc.GetUserbyName)
	ec.GET("/users/:id", uc.GetUserbyID)
	ec.POST("/verify", uc.VerifyUser)
	ec.POST("/users", uc.InsertUser)
	ec.POST("/ret", controllers.PaymentStatus)

	ec.POST("/books", bc.InsertBook)

	ecAuth := ec.Group("")
	ecAuth.Use(middleware.JWT([]byte("rahasia")))
	ecAuth.DELETE("/users/:id", uc.DeleteUser)
}
