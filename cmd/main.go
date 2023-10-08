package main

import (
	"net/http"

	_ "github.com/gitkuldp/wedding-invitation-api/docs"
	customMiddleware "github.com/gitkuldp/wedding-invitation-api/internal/api/middleware"
	v1 "github.com/gitkuldp/wedding-invitation-api/internal/api/v1/routes"

	"github.com/gitkuldp/wedding-invitation-api/internal/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title						Invitation API Documentation
// @version					1.0
// @description				This is an API documentation of Invitation.
//
// @BasePath					/api/v1/erp
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							Authorization
// @name						Bearer
func main() {
	app := db.App()
	env := app.Env
	echo := echo.New()
	echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", "localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	echo.Use(customMiddleware.Authenticator)
	echo.Validator = db.NewValidator()
	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())
	echo.GET("/swagger/*", echoSwagger.WrapHandler)

	v1.Setup(&app, echo)
	err := echo.Start("https://" + env.ServerAddress + ":8080")
	logrus.Error(err)
}
