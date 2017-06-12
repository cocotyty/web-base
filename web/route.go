package web

import (
	"time"

	"github.com/cocotyty/summer"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/view"
)

func init() {
	app := iris.New()
	route := newRoute(app)
	summer.Add("web", route)
}

type Route struct {
	Debug  bool   `sm:"#.debug"`
	Listen string `sm:"#.listen"`

	app *iris.Application
}

func newRoute(app *iris.Application) *Route {
	return &Router{app: app}
}

func (r *Route) Init(app *iris.Application) {
	if r.Debug {
		context.StaticCacheDuration = 1 * time.Second
	}

	r.app.Use(recover.New())

	if r.Debug {
		r.app.UseGlobal(func(ctx context.Context) {
			ctx.Gzip(true)
			ctx.Next()
		})
	}

	r.app.AttachView(view.HTML("./templates", ".html").Delims("$$"))
}

func (r *Route) Config() {
	r.app.Get("/", func(ctx context.Context) {
		ctx.ViewData("Message", "Hello World")
		ctx.View("index.html")
	})
}
func (r *Route) Ready() {
	r.Config()
}

func (r *Route) Start() error {
	//
	// TODO: load configuration from the existing
	// toml files at ./conf.
	// Read more here: https://github.com/kataras/iris/blob/master/_examples/beginner/configuration/from-toml-file
	//
	return r.app.Run(iris.Addr(r.Listen)) // gracefully shutdowns by-default now
}
