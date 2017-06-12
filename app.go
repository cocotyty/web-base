package main

import (
	"os"

	"github.com/cocotyty/summer"
	_ "github.com/cocotyty/web-base/provider"
	"github.com/cocotyty/web-base/web"
)

func main() {
	mode := os.Getenv("app_mode")
	if mode == "" {
		mode = "dev"
	}

	summer.TomlFile("./conf/" + mode + ".toml")
	summer.Start()

	if err := summer.GetStoneWithName("web").(*web.Route).Start(); err != nil {
		panic(err)
	}
}
