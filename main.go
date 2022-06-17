package main

import (
	"ziweiBlog/config"
	"ziweiBlog/server"
)

func main() {
	server.App.Start(config.Cfg.System.Host, config.Cfg.System.Port)
}
