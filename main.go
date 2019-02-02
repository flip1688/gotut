package main

import (
	"flag"
	"fmt"
	"gotut/database"
	"gotut/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type response struct {
	Status  string `json:"status"`
	Message string
}
type loginRequest struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Repsonse response `json:"response"`
}

type CreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	userSvc service.Service
)

func main() {
	port := flag.Int("port", 1323, "Port number for http server")
	fmt.Printf("%v\n", *port)

	res := response{
		Status:  "OK",
		Message: "Naja",
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, res)
	})
	e.POST("/", func(c echo.Context) error {
		var postBody loginRequest
		if err := c.Bind(&postBody); err != nil {
			return c.String(http.StatusBadRequest, "Internal error")
		}
		foundUser, err := userSvc.VerifyUserPassword(postBody.Username, postBody.Password)
		if err == nil {
			return c.JSON(http.StatusOK, foundUser)
		}
		if err == service.ErrInvalidLogin {
			return c.String(http.StatusUnauthorized, err.Error())
		}
		log.Printf("error when login message: %v", err)
		return c.String(http.StatusInternalServerError, err.Error())
	})
	userPath := e.Group("/user")
	userPath.POST("", createUser)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(*port)))
}

func createUser(c echo.Context) error {
	var reqBody CreateUser
	if err := c.Bind(&reqBody); err != nil {
		return c.String(http.StatusBadRequest, "Internal error")
	}
	user := database.User{
		Username: reqBody.Username,
		Password: reqBody.Password,
	}

	if err := userSvc.CreateUser(&user); err != nil {
		log.Printf("error when create user message: %v", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}
