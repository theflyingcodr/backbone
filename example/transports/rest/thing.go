package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"github.com/theflyingcodr/things"
)

type thing struct {
	svc things.ThingService
}

// NewEndpoint will setup and return a new thing handler.
func NewEndpoint(svc things.ThingService) *thing {
	return &thing{svc: svc}
}

// RegisterRoutes is called in main and takes an echo group which will wire up
// the endpoints to the group.
func (e *thing) RegisterRoutes(g *echo.Group) {
	g.GET(routeThing, e.thing)
	g.GET(routeThings, e.things)
}

func (e *thing) thing(c echo.Context) error {
	var args things.ThingArgs
	if err := c.Bind(&args); err != nil {
		return errors.Wrap(err, "failed to parse arguments")
	}
	resp, err := e.svc.Thing(c.Request().Context(), args)
	if err != nil {
		return errors.WithStack(err)
	}
	return c.JSON(http.StatusOK, resp)
}

func (e *thing) things(c echo.Context) error {
	resp, err := e.svc.Things(c.Request().Context())
	if err != nil {
		return errors.WithStack(err)
	}
	return c.JSON(http.StatusOK, resp)
}

func (e *thing) create(c echo.Context) error {
	var req things.ThingCreate
	if err := c.Bind(&req); err != nil {
		return errors.Wrap(err, "failed to parse request")
	}
	resp, err := e.svc.Things(c.Request().Context())
	if err != nil {
		return errors.WithStack(err)
	}
	return c.JSON(http.StatusCreated, resp)
}

func (e *thing) update(c echo.Context) error {
	var args things.ThingArgs
	if err := c.Bind(&args); err != nil {
		return errors.Wrap(err, "failed to parse arguments")
	}
	var req things.ThingUpdate
	if err := c.Bind(&req); err != nil {
		return errors.Wrap(err, "failed to parse request")
	}
	resp, err := e.svc.Update(c.Request().Context(), args, req)
	if err != nil {
		return errors.WithStack(err)
	}
	return c.JSON(http.StatusOK, resp)
}

func (e *thing) delete(c echo.Context) error {
	var args things.ThingArgs
	if err := c.Bind(&args); err != nil {
		return errors.Wrap(err, "failed to parse arguments")
	}
	if err := e.svc.Delete(c.Request().Context(), args); err != nil {
		return errors.WithStack(err)
	}
	return c.NoContent(http.StatusNoContent)
}
