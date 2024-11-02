package router

import (
	"net/http"
	"os"
	"github.com/Vasudev-2308/gostudy/intenal/config"
)

const (
	StudentDB = "students"
	TeacherDB = "teachers"
)

func StartRouter(cfg config.Config) {

	router := *http.NewServeMux()
	//All End Points of
	initRoutes(&router, cfg)

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: &router,
	}

	doneChannel := make(chan os.Signal, 1)
	StartServer(&server, doneChannel)
}
