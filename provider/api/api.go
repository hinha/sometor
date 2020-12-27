package api

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hinha/sometor/provider"
	"github.com/hinha/sometor/provider/api/command"
	"github.com/hinha/sometor/provider/api/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
)

// API ...
type API struct {
	engine *echo.Echo
	port   int
}

func Fabricate(givenPort int) *API {
	return &API{
		engine: echo.New(),
		port:   givenPort,
	}
}

// FabricateCommand insert api related command
func (a *API) FabricateCommand(cmd provider.Command) {
	cmd.InjectCommand(
		command.NewRun(a),
	)
}

// InjectAPI inject new API into api_box
func (a *API) InjectAPI(handler provider.APIHandler) {
	a.engine.Add(handler.Method(), handler.Path(), func(context echo.Context) error {
		req := context.Request()
		if reqID := req.Header.Get("X-Request-ID"); reqID != "" {
			context.Set("request-id", reqID)
		} else {
			context.Set("request-id", uuid.New().String())
		}

		if userID := req.Header.Get("Resource-Owner-ID"); userID != "" {
			convertedUserID, err := strconv.Atoi(userID)
			if err == nil {
				context.Set("user-id", convertedUserID)
			}
		}

		handler.Handle(context)

		return nil
	})
}

func (a *API) Run() error {
	a.engine.Use(middleware.Logger())
	a.engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderAccessControlAllowOrigin,
		},
	}))
	a.InjectAPI(handler.NewHealth())
	return a.engine.Start(fmt.Sprintf(":%d", a.port))
}

// Shutdown api engine
func (a *API) Shutdown(ctx context.Context) error {
	return a.engine.Shutdown(ctx)
}
