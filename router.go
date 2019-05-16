package openapi3

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo"
)

const DefaultOpenAPIVersion = "3.0.0"

type Echo struct {
	*echo.Echo
	OpenAPI *openapi3.Swagger
}

type EchoGroup struct {
	*echo.Group
	OpenAPI *openapi3.Swagger
	prefix  string
}

func NewEcho() *Echo {
	return &Echo{Echo: echo.New(), OpenAPI: &openapi3.Swagger{OpenAPI: DefaultOpenAPIVersion}}
}

func (r *Echo) Group(path string, mw ...echo.MiddlewareFunc) *EchoGroup {
	return &EchoGroup{r.Echo.Group(path, mw...), r.OpenAPI, path}
}

// SubGroup can be used to create a Group from a Group.
// This method is simply named Group in the Echo library but has to
// be renamed to not clash with the embedded Group object.
func (r *EchoGroup) SubGroup(path string, mw ...echo.MiddlewareFunc) *EchoGroup {
	return &EchoGroup{r.Group.Group(path, mw...), r.OpenAPI, r.prefix + path}
}

func (e *Echo) GET(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	e.Echo.GET(path, handler)
	if op != nil {
		e.OpenAPI.AddOperation(path, "GET", op)
	}
}
func (e *EchoGroup) GET(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	e.Group.GET(path, handler)
	if op != nil {
		e.OpenAPI.AddOperation(e.prefix+path, "GET", op)
	}
}

func (e *Echo) POST(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	e.Echo.POST(path, handler)
	if op != nil {
		e.OpenAPI.AddOperation(path, "POST", op)
	}
}
func (e *EchoGroup) POST(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	e.Group.POST(path, handler)
	if op != nil {
		e.OpenAPI.AddOperation(e.prefix+path, "POST", op)
	}
}

func (e *Echo) PUT(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	e.Echo.PUT(path, handler)
	if op != nil {
		e.OpenAPI.AddOperation(path, "PUT", op)
	}
}
func (e *EchoGroup) PUT(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	e.Group.PUT(path, handler)
	if op != nil {
		e.OpenAPI.AddOperation(e.prefix+path, "PUT", op)
	}
}

func (e *Echo) PATCH(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	e.Echo.PATCH(path, handler)
	if op != nil {
		e.OpenAPI.AddOperation(path, "PATCH", op)
	}
}
func (e *EchoGroup) PATCH(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	e.Group.PATCH(path, handler)
	if op != nil {
		e.OpenAPI.AddOperation(e.prefix+path, "PATCH", op)
	}
}

func (e *Echo) DELETE(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	e.Echo.DELETE(path, handler)
	if op != nil {
		e.OpenAPI.AddOperation(path, "DELETE", op)
	}
}
func (e *EchoGroup) DELETE(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	e.Group.DELETE(path, handler)
	if op != nil {
		e.OpenAPI.AddOperation(e.prefix+path, "DELETE", op)
	}
}

func (e *Echo) CONNECT(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	e.Echo.CONNECT(path, handler)
	if op != nil {
		e.OpenAPI.AddOperation(path, "CONNECT", op)
	}
}
func (e *EchoGroup) CONNECT(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	e.Group.CONNECT(path, handler)
	if op != nil {
		e.OpenAPI.AddOperation(e.prefix+path, "CONNECT", op)
	}
}

func (e *Echo) HEAD(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	e.Echo.HEAD(path, handler)
	if op != nil {
		e.OpenAPI.AddOperation(path, "HEAD", op)
	}
}
func (e *EchoGroup) HEAD(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	e.Group.HEAD(path, handler)
	if op != nil {
		e.OpenAPI.AddOperation(e.prefix+path, "HEAD", op)
	}
}

func (e *Echo) OPTIONS(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	e.Echo.OPTIONS(path, handler)
	if op != nil {
		e.OpenAPI.AddOperation(path, "OPTIONS", op)
	}
}
func (e *EchoGroup) OPTIONS(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	e.Group.OPTIONS(path, handler)
	if op != nil {
		e.OpenAPI.AddOperation(e.prefix+path, "OPTIONS", op)
	}
}

func (e *Echo) TRACE(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	e.Echo.TRACE(path, handler)
	if op != nil {
		e.OpenAPI.AddOperation(path, "TRACE", op)
	}
}
func (e *EchoGroup) TRACE(path string, handler echo.HandlerFunc, op *openapi3.Operation) {
	e.Group.TRACE(path, handler)
	if op != nil {
		e.OpenAPI.AddOperation(e.prefix+path, "TRACE", op)
	}
}

func (e *Echo) Any(path, swaggerMethod string, handler echo.HandlerFunc, op *openapi3.Operation) {
	e.Echo.Any(path, handler)
	if op != nil {
		e.OpenAPI.AddOperation(path, swaggerMethod, op)
	}
}
func (e *EchoGroup) Any(path, swaggerMethod string, handler echo.HandlerFunc, op *openapi3.Operation) {
	e.Group.Any(path, handler)
	if op != nil {
		e.OpenAPI.AddOperation(e.prefix+path, swaggerMethod, op)
	}
}
