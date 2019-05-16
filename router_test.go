package openapi3

import (
	"testing"

	"github.com/labstack/echo"
)

func TestRouter(t *testing.T) {
	r := NewRouter()
	r.GET("/", func(ctx echo.Context) error {
		return nil
	}, nil)
	numRoutes := len(r.Echo.Routes())
	if numRoutes != 1 {
		t.Errorf("expected one registered route on echo, got %d", numRoutes)
	}
}
