package controllers

import (
	"net/http"
	"project/mock_api/middlewares"
	"project/mock_api/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userModel models.UserModel
}

func NewController(um models.UserModel) *UserController {
	return &UserController{um}
}

func AuthorizeAdmin(ec echo.Context) bool {
	_, role := middlewares.ExtractTokenUser(ec)
	return role == "admin"
}

func (c *UserController) HelloWorld(ec echo.Context) error {
	return ec.JSON(http.StatusOK, "HELLO WORLD")
}

func (c *UserController) GetAllUser(ec echo.Context) error {
	users, err := c.userModel.GetAll()

	checkAuthorize := AuthorizeAdmin(ec)
	if !checkAuthorize {
		return ec.JSON(http.StatusUnauthorized, map[string]interface{}{
			"success": false,
			"code":    401,
			"message": "Unauthorized Error",
		})
	}

	if err != nil {
		return ec.JSON(http.StatusInternalServerError, "cannot get users")
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all user",
		"data":    users,
	})
}

func (c *UserController) GetUserbyID(ec echo.Context) error {
	name, err := strconv.Atoi(ec.Param("id"))

	user, err := c.userModel.GetByID(name)

	if err != nil {
		return ec.JSON(http.StatusInternalServerError, "cannot get users")
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get user",
		"data":    user,
	})
}

func (c *UserController) GetUserbyName(ec echo.Context) error {
	name := ec.QueryParam("name")

	user, err := c.userModel.GetByName(name)

	if err != nil {
		return ec.JSON(http.StatusInternalServerError, "cannot get users")
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get user",
		"data":    user,
	})
}

func (c *UserController) InsertUser(ec echo.Context) error {
	user := models.User{}
	ec.Bind(&user)
	user, err := c.userModel.Insert(user)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, "invalid user data")
	}
	return ec.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success add user",
		"data":    user,
	})
}

func (c *UserController) DeleteUser(ec echo.Context) error {
	id, err := strconv.Atoi(ec.Param("id"))

	if err != nil {
		return ec.JSON(http.StatusBadRequest, "invalid id")
	}

	user, err := c.userModel.Delete(id)

	if user.ID == 0 {
		return ec.JSON(http.StatusNotFound, "data not found")
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete user",
		"data":    user,
	})
}

func (c *UserController) VerifyUser(ec echo.Context) error {
	user := models.User{}
	ec.Bind(&user)
	user, err := c.userModel.GetByEmailAndPassword(user.Email, user.Password)

	if err != nil {
		return ec.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
		})
	}

	token, err := middlewares.CreateToken(int(user.ID), "rahasia")
	user.Token = token

	user, err = c.userModel.Update(user, int(user.ID))
	if err != nil {
		return ec.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
		})
	}

	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot add token",
		})
	}
	return ec.JSON(http.StatusOK, map[string]interface{}{
		"message": "user have been verified",
		"data":    user,
	})

}

// func (c *UserController) Login(ec echo.Context) error {
// 	user := models.User{}
// 	ec.Bind(&user)

// 	user, err := c.userModel.GetByEmailAndPassword(user.Email, user.Password)

// 	if err != nil {
// 		return ec.JSON(http.StatusNotFound, "user not found")
// 	}
// }
