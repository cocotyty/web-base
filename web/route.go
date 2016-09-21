package web

import (
	"github.com/cocotyty/summer"
	"github.com/kataras/iris"
	"time"
	"github.com/iris-contrib/middleware/recovery"
	"github.com/kataras/iris/config"
	"github.com/iris-contrib/template/html"
	"github.com/iris-contrib/graceful"
)

func init() {
	summer.Add("web",&Route{})
}

type Route struct {
	Debug  bool `sm:"#.debug"`
	Listen string `sm:"#.listen"`
}

func (r *Route)Init() {
	if r.Debug {
		config.StaticCacheDuration = 1 * time.Second
	}
	iris.Config.IsDevelopment = r.Debug
	iris.Config.Gzip = !r.Debug
	iris.UseTemplate(html.New(html.Config{
		Left:        "$$",
		Right:       "$$",
		Layout:      "",
		Funcs:       make(map[string]interface{}, 0),
		LayoutFuncs: make(map[string]interface{}, 0),
	}))
	iris.Use(recovery.New(iris.Logger))
}
func (r *Route)Config() {
	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("index.html", "Hello World")
	})
}
func (r *Route)Ready() {
	r.Config()
}

func (r *Route)Start() {
	if r.Debug {
		iris.Listen(r.Listen)
	} else {
		graceful.Run(r.Listen, 5 * time.Second, iris.Default)
	}
}