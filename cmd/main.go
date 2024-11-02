package main

import (
	router "github.com/Vasudev-2308/gostudy/cmd/src"
	"github.com/Vasudev-2308/gostudy/intenal/config"
)

func main() {
	// Load Config
	cfg := config.Load()
	// Starting Router and Server {startServer is embedded inside startROuter}
	router.StartRouter(*cfg)

}
