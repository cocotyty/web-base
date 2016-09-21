package main

import (
	_ "github.com/cocotyty/web-base/provider"
	"github.com/cocotyty/summer"
	"github.com/cocotyty/web-base/web"
	"os"
)

func main() {
	mode := os.Getenv("app_mode")
	if mode == "" {
		mode = "dev"
	}
	summer.TomlFile("./conf/" + mode + ".toml")
	summer.Start()
	summer.GetStoneWithName("web").(*web.Route).Start()
}
