package provider

import (
	"context"
	"github.com/labstack/echo/v4"
	"mime/multipart"
	"net/http"
	"net/url"
)

// APIContext used by API handler to modify it's request
type APIContext interface {
	// Response returns `*Response`.
	Response() *echo.Response

	// Request returns `*http.Request`.
	Request() *http.Request

	// RealIP returns the client's network address based on `X-Forwarded-For`
	// or `X-Real-IP` request header.
	// The behavior can be configured using `Echo#IPExtractor`.
	RealIP() string

	// Path returns the registered path for the handler.
	Path() string

	// Param returns path parameter by name.
	Param(name string) string

	// ParamNames returns path parameter names.
	ParamNames() []string

	// ParamValues returns path parameter values.
	ParamValues() []string

	// QueryParam returns the query param for the provided name.
	QueryParam(name string) string

	// QueryParams returns the query parameters as `url.Values`.
	QueryParams() url.Values

	// QueryString returns the URL query string.
	QueryString() string

	// FormFile returns the multipart form file for the provided name.
	FormFile(name string) (*multipart.FileHeader, error)

	// Cookie returns the named cookie provided in the request.
	Cookie(name string) (*http.Cookie, error)

	// SetCookie adds a `Set-Cookie` header in HTTP response.
	SetCookie(cookie *http.Cookie)

	// Cookies returns the HTTP cookies sent with the request.
	Cookies() []*http.Cookie

	// Get retrieves data from the context.
	Get(key string) interface{}

	// Set saves data in the context.
	Set(key string, val interface{})

	// Bind binds the request body into provided type `i`. The default binder
	// does it based on Content-Type header.
	Bind(i interface{}) error

	// JSON sends a JSON response with status code.
	JSON(code int, i interface{}) error

	// NoContent sends a response with no body and a status code.
	NoContent(code int) error

	// JSONPretty sends a pretty-print JSON with status code.
	JSONPretty(code int, i interface{}, indent string) error

	// JSONBlob sends a JSON blob response with status code.
	JSONBlob(code int, b []byte) error

	// JSONP sends a JSONP response with status code. It uses `callback` to construct
	// the JSONP payload.
	JSONP(code int, callback string, i interface{}) error
}

// APIHandler handling api request from client
type APIHandler interface {
	Handle(context APIContext)
	Method() string
	Path() string
}

// APIEngine ...
type APIEngine interface {
	Run() error
	InjectAPI(handler APIHandler)
	Shutdown(ctx context.Context) error
}
