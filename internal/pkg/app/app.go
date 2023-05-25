package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"

	"github.com/Valley-Craft/VCBE/internal/app/endpoint"
	"github.com/Valley-Craft/VCBE/internal/app/mw"
	"github.com/Valley-Craft/VCBE/internal/app/service"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(middleware.CORS())

	a.echo.Use(mw.RoleCheck)

	a.echo.GET("api/v1/players", a.e.PlayersEndPoint)
	a.echo.POST("api/v1/form", a.e.FormEndPoint)

	return a, nil
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func (a *App) Run() error {

	port, _ := os.LookupEnv("PORT")

	err := a.echo.Start(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
