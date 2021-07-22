package router

import (
	"attendance-server/database"
	"fmt"
	"github.com/labstack/echo/middleware"
	"net/http"

	"github.com/labstack/echo"
)

type APIServer struct {
	E *echo.Echo
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

const (
	StatusNotFound = "NOT_FOUND_SERVICE"
)

func mysqlMiddleware(mysqlDatabase *database.MysqlClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("mysql", mysqlDatabase)
			return next(c)
		}
	}
}

func New(config *ServerConfig, mysqlDatabase *database.MysqlClient) *APIServer {
	server := new(APIServer)
	serverUrl := fmt.Sprintf("%s:%d", config.Host, config.Port)

	echo.NotFoundHandler = func(c echo.Context) error {
		return c.String(http.StatusNotFound, StatusNotFound)
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(mysqlMiddleware(mysqlDatabase))

	e.Logger.Fatal(e.Start(serverUrl))

	return server
}
