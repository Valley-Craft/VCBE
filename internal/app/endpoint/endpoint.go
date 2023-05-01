package endpoint

import (
	"github.com/labstack/echo/v4"
	"io"
	"io/ioutil"
	"net/http"
)

type Service interface {
	Players() string
	Form(string) bool
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
	data := e.s.Players()

	err := ctx.String(http.StatusOK, data)
	if err != nil {
		return err
	}

	return nil
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
