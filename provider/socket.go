package provider

import (
	"context"
	"github.com/hinha/sometor/entity"
	"github.com/labstack/echo/v4"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

type SocketContext interface {
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

	// Validate validates provided `i`. It is usually called after `Context#Bind()`.
	// Validator must be registered using `Echo#Validator`.
	Validate(i interface{}) error

	// Render renders a template with data and sends a text/html response with status
	// code. Renderer must be registered using `Echo.Renderer`.
	Render(code int, name string, data interface{}) error

	// JSON sends a JSON response with status code.
	JSON(code int, i interface{}) error

	// JSONPretty sends a pretty-print JSON with status code.
	JSONPretty(code int, i interface{}, indent string) error

	// JSONBlob sends a JSON blob response with status code.
	JSONBlob(code int, b []byte) error

	// JSONP sends a JSONP response with status code. It uses `callback` to construct
	// the JSONP payload.
	JSONP(code int, callback string, i interface{}) error

	// Blob sends a blob response with status code and content type.
	Blob(code int, contentType string, b []byte) error

	// Stream sends a streaming response with status code and content type.
	Stream(code int, contentType string, r io.Reader) error

	// File sends a response with the content of the file.
	File(file string) error

	// Attachment sends a response as attachment, prompting client to save the
	// file.
	Attachment(file string, name string) error

	// Inline sends a response as inline, opening the file in the browser.
	Inline(file string, name string) error

	// NoContent sends a response with no body and a status code.
	NoContent(code int) error

	// Redirect redirects the request to a provided URL with status code.
	Redirect(code int, url string) error

	// Error invokes the registered HTTP error handler. Generally used by middleware.
	Error(err error)

	// Logger returns the `Logger` instance.
	Logger() echo.Logger
}

type SocketHandler interface {
	Handle(context SocketContext)
	Method() string
	Path() string
}

type SocketEngine interface {
	Run() error
	InjectSocket(handler SocketHandler)
	Shutdown(ctx context.Context) error
}

type SocketMedia interface {
	FileReader(ctx context.Context, lastMod time.Time, media string, FileUser string) ([]byte, time.Time, *entity.ApplicationError)
	Writer(ctx context.Context, ID string, FileUser string) // deprecate
	UserValid(ctx context.Context, ID, keyword, media string) (entity.UserAccountSelectable, *entity.ApplicationError)
}
