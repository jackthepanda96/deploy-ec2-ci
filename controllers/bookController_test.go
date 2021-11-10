package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	configs "project/mock_api/config"

	"testing"

	"project/mock_api/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func SetUp() {
	config := configs.GetConfig()
	db := configs.MysqlConnect(config)

	userModel := models.NewUserModel(db)
	// bookModel := models.NewBookModel(db)

	// newBook := models.Book{}
	// newBook.Title = "Buku Pintar"
	// _, err := bookModel.Insert(newBook)

	newUser := models.User{}
	newUser.Name = "Jerry"
	newUser.Email = "Jerry@jerry.com"
	newUser.Gender = "L"
	newUser.Password = "rahasia"

	_, err := userModel.Insert(newUser)

	if err != nil {
		fmt.Println(err)
	}
}

func TestLogin(t *testing.T) {
	config := configs.GetConfig()
	db := configs.MysqlConnect(config)
	userModel := models.NewUserModel(db)
	userController := NewController(userModel)

	e := echo.New()

	e.POST("/verify", userController.VerifyUser)

	reqBody, _ := json.Marshal(map[string]string{
		"email":    "jerry@Jerry.com",
		"password": "rahasia",
	})

	req := httptest.NewRequest(echo.POST, "/verify", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	type Response struct {
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
	var response Response
	resBody := rec.Body.String()
	json.Unmarshal([]byte(resBody), &response)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "user have been verified", response.Message)

	reqBody2, _ := json.Marshal(map[string]string{
		"email":    "jerry@Jerry.com",
		"password": "",
	})

	req2 := httptest.NewRequest(echo.POST, "/verify", bytes.NewBuffer(reqBody2))
	req.Header.Set("Content-Type", "application/json")
	rec2 := httptest.NewRecorder()
	e.ServeHTTP(rec2, req2)

	var response2 Response
	resBody2 := rec2.Body.String()
	json.Unmarshal([]byte(resBody2), &response)

	assert.Equal(t, http.StatusNotFound, rec2.Code)
	assert.Empty(t, response2.Data)
	assert.Equal(t, "user not found", response2.Message)

	e.Server.Close()
}

func TestGetAllUser(t *testing.T) {
	// SetUp()
	// Initialize configuration
	// Init echo
	// Configure echo route

	// Access each endpoint using e.ServeHTTP
	config := configs.GetConfig()
	db := configs.MysqlConnect(config)
	userModel := models.NewUserModel(db)
	userController := NewController(userModel)

	e := echo.New()

	e.POST("/verify", userController.VerifyUser)
	e.GET("/users", userController.GetAllUser, middleware.JWT([]byte("rahasia")))

	reqBody, _ := json.Marshal(map[string]string{
		"email":    "jerry@Jerry.com",
		"password": "rahasia",
	})

	req := httptest.NewRequest(echo.POST, "/verify", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	type Response struct {
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
	var response Response
	resBody := rec.Body.String()
	json.Unmarshal([]byte(resBody), &response)

	assert.Equal(t, http.StatusOK, rec.Code)

	req2 := httptest.NewRequest(echo.GET, "/users", nil)
	rec2 := httptest.NewRecorder()
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("Authorization", fmt.Sprint("Bearer ", response.Data.(map[string]interface{})["token"]))
	e.ServeHTTP(rec2, req2)

	var response2 Response
	resBody2 := rec2.Body.String()
	json.Unmarshal([]byte(resBody2), &response2)

	assert.Equal(t, http.StatusOK, rec2.Code)
	assert.Equal(t, "Success get all user", response2.Message)
	e.Server.Close()
}

func TestGetIDUser(t *testing.T) {
	// Access each endpoint using e.ServeHTTP
	config := configs.GetConfig()
	db := configs.MysqlConnect(config)
	userModel := models.NewUserModel(db)
	userController := NewController(userModel)

	e := echo.New()

	e.POST("/verify", userController.VerifyUser)
	e.GET("/users/:id", userController.GetUserbyID, middleware.JWT([]byte("rahasia")))

	reqBody, _ := json.Marshal(map[string]string{
		"email":    "jerry@Jerry.com",
		"password": "rahasia",
	})

	req := httptest.NewRequest(echo.POST, "/verify", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	type Response struct {
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
	var response Response
	resBody := rec.Body.String()
	json.Unmarshal([]byte(resBody), &response)

	assert.Equal(t, http.StatusOK, rec.Code)

	req2 := httptest.NewRequest(echo.GET, "/users/6", nil)
	rec2 := httptest.NewRecorder()
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("Authorization", fmt.Sprint("Bearer ", response.Data.(map[string]interface{})["token"]))
	e.ServeHTTP(rec2, req2)

	var response2 Response
	resBody2 := rec2.Body.String()
	json.Unmarshal([]byte(resBody2), &response2)

	assert.Equal(t, http.StatusOK, rec2.Code)
	assert.Equal(t, "Success get user", response2.Message)
	e.Server.Close()
}
