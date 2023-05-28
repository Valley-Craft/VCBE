package endpoint

import (
	"github.com/Valley-Craft/VCBE/internal/app/service"
	"github.com/labstack/echo/v4"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Service interface {
	Players() ([]service.JSONPlayer, error)
	Form(string) bool
	Donate(string) bool
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) PlayersEndPoint(ctx echo.Context) error {
	data, _ := e.s.Players()

	err := ctx.JSON(http.StatusOK, data)
	if err != nil {
		return err
	}

	return nil
}

func (e *Endpoint) DonateEndPoint(ctx echo.Context) error {

	xKey := ctx.Request().Header.Get("X-Key")

	key, _ := os.LookupEnv("X_KEY_DONATE")

	if xKey != key {
		err := ctx.String(http.StatusNotAcceptable, "No access")
		return err
	}

	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(ctx.Request().Body)

	data := e.s.Donate(string(body))

	if data {
		err = ctx.String(http.StatusAccepted, "true")
		if err != nil {
			return err
		}
	}

	err = ctx.String(http.StatusBadRequest, "false")
	if err != nil {
		return err
	}

	return err
}

func (e *Endpoint) FormEndPoint(ctx echo.Context) error {

	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(ctx.Request().Body)

	data := e.s.Form(string(body))

	if data {
		err = ctx.String(http.StatusAccepted, "")
		if err != nil {
			return err
		}
	}

	err = ctx.String(http.StatusBadRequest, "")
	if err != nil {
		return err
	}

	return nil
}
