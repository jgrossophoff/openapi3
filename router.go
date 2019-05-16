package openapi3

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo"
)

type Router struct {
	*echo.Echo
	OpenAPI *openapi3.Swagger
}

func NewRouter() *Router {
	return &Router{Echo: echo.New(), OpenAPI: &openapi3.Swagger{}}
}

func (r *Router) Group(path string, mw ...echo.MiddlewareFunc) *RouteGroup {
	return &RouteGroup{r.Echo.Group(path, mw...), r.OpenAPI}
}

type RouteGroup struct {
	*echo.Group
	OpenAPI *openapi3.Swagger
}

func (s *Router) GET(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	s.Echo.GET(path, handler)
	if op != nil {
		s.OpenAPI.AddOperation(path, "GET", op)
	}
}
func (s *RouteGroup) GET(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	s.Group.GET(path, handler)
	if op != nil {
		s.OpenAPI.AddOperation(path, "GET", op)
	}
}

func (s *Router) POST(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	s.Echo.POST(path, handler)
	if op != nil {
		s.OpenAPI.AddOperation(path, "POST", op)
	}
}
func (s *RouteGroup) POST(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	s.Group.POST(path, handler)
	if op != nil {
		s.OpenAPI.AddOperation(path, "POST", op)
	}
}

func (s *Router) PUT(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	s.Echo.PUT(path, handler)
	if op != nil {
		s.OpenAPI.AddOperation(path, "PUT", op)
	}
}
func (s *RouteGroup) PUT(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	s.Group.PUT(path, handler)
	if op != nil {
		s.OpenAPI.AddOperation(path, "PUT", op)
	}
}

func (s *Router) PATCH(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	s.Echo.PATCH(path, handler)
	if op != nil {
		s.OpenAPI.AddOperation(path, "PATCH", op)
	}
}
func (s *RouteGroup) PATCH(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	s.Group.PATCH(path, handler)
	if op != nil {
		s.OpenAPI.AddOperation(path, "PATCH", op)
	}
}

func (s *Router) DELETE(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	s.Echo.DELETE(path, handler)
	if op != nil {
		s.OpenAPI.AddOperation(path, "DELETE", op)
	}
}
func (s *RouteGroup) DELETE(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	s.Group.DELETE(path, handler)
	if op != nil {
		s.OpenAPI.AddOperation(path, "DELETE", op)
	}
}

func (s *Router) CONNECT(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	s.Echo.CONNECT(path, handler)
	if op != nil {
		s.OpenAPI.AddOperation(path, "CONNECT", op)
	}
}
func (s *RouteGroup) CONNECT(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	s.Group.CONNECT(path, handler)
	if op != nil {
		s.OpenAPI.AddOperation(path, "CONNECT", op)
	}
}

func (s *Router) HEAD(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	s.Echo.HEAD(path, handler)
	if op != nil {
		s.OpenAPI.AddOperation(path, "HEAD", op)
	}
}
func (s *RouteGroup) HEAD(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	s.Group.HEAD(path, handler)
	if op != nil {
		s.OpenAPI.AddOperation(path, "HEAD", op)
	}
}

func (s *Router) OPTIONS(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	s.Echo.OPTIONS(path, handler)
	if op != nil {
		s.OpenAPI.AddOperation(path, "OPTIONS", op)
	}
}
func (s *RouteGroup) OPTIONS(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	s.Group.OPTIONS(path, handler)
	if op != nil {
		s.OpenAPI.AddOperation(path, "OPTIONS", op)
	}
}

func (s *Router) TRACE(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	s.Echo.TRACE(path, handler)
	if op != nil {
		s.OpenAPI.AddOperation(path, "TRACE", op)
	}
}
func (s *RouteGroup) TRACE(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	s.Group.TRACE(path, handler)
	if op != nil {
		s.OpenAPI.AddOperation(path, "TRACE", op)
	}
}

func (s *Router) Any(path, swaggerMethod string, handler echo.HandlerFunc, op *openapi3.Operation) {
	s.Echo.Any(path, handler)
	if op != nil {
		s.OpenAPI.AddOperation(path, swaggerMethod, op)
	}
}
func (s *RouteGroup) Any(path, swaggerMethod string, handler echo.HandlerFunc, op *openapi3.Operation) {
	s.Group.Any(path, handler)
	if op != nil {
		s.OpenAPI.AddOperation(path, swaggerMethod, op)
	}
}
