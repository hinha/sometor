package socket

import (
	"context"
	"fmt"
	"github.com/hinha/sometor/provider"
	"github.com/hinha/sometor/provider/socket/command"
	"github.com/hinha/sometor/provider/socket/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
)

type Socket struct {
	engine *echo.Echo
	port   int
}

func Fabricate(givenPort int) *Socket {
	return &Socket{engine: echo.New(), port: givenPort}
}

// FabricateCommand insert api related command
func (a *Socket) FabricateCommand(cmd provider.Command) {
	cmd.InjectCommand(
		command.NewRunSocket(a),
	)
}

func (a *Socket) Run() error {
	a.engine.Use(middleware.Logger())
	//a.engine.Static("/", "assets")
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("assets/*.html")),
	}
	a.engine.Renderer = renderer

	a.InjectAPI(handler.NewPing())
	a.InjectAPI(handler.NewPingWeb())
	return a.engine.Start(fmt.Sprintf(":%d", a.port))
}

func (a *Socket) InjectAPI(handler provider.SocketHandler) {
	a.engine.Add(handler.Method(), handler.Path(), func(context echo.Context) error {
		handler.Handle(context)
		return nil
	})
}

func (a *Socket) Shutdown(ctx context.Context) error {
	return a.engine.Close()
}

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
