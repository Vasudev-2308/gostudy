package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Vasudev-2308/gostudy/intenal/config"
	"github.com/Vasudev-2308/gostudy/intenal/http/handlers/User"
	"github.com/Vasudev-2308/gostudy/intenal/storage/sqlite"
)

func startServer(server *http.Server, doneChannel chan os.Signal) {
	// Setup server and Gracefully Chose it Instead of Abruptly Shutting it Down
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			slog.Error(err.Error())
		}
	}()

	signal.Notify(doneChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-doneChannel

	slog.Info("Shutting Down")

	ctx, cancel := context.WithTimeout(context.Background(), 5)
	defer cancel()

	error := server.Shutdown(ctx)

	if error != nil {
		slog.Error("Failed to Shutdown", slog.String("error: ", error.Error()))
	}
	slog.Info("Server Showdown Successfully")
}

func startRouter(cfg config.Config) {
	// Setup Router
	// Setup Database
	student, err1 := sqlite.NewStudent(&cfg)
	teacher, err2 := sqlite.NewTeacher(&cfg)

	if err1 != nil || err2 != nil {
		slog.Info("Not able to Initiate DB", slog.String("%s", err1.Error()), slog.String("%s", err2.Error()))
	}

	slog.Info("Storage Initated", slog.String("env", cfg.StoragePath))
	router := http.NewServeMux()
	router.HandleFunc("GET /api/student/{id}", User.GetUser(student, "students"))
	router.HandleFunc("POST /api/create-student", User.AddUser(student, "students"))
	router.HandleFunc("GET /api/teachers", User.GetUser(teacher, "teachers"))
	router.HandleFunc("POST /api/create-teachers", User.AddUser(teacher, "teachers"))

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	doneChannel := make(chan os.Signal, 1)
	startServer(&server, doneChannel)
}

func main() {
	// Load Config
	cfg := config.MustLoad()
	// Starting Router and Server {startServer is embedded inside startROuter}
	startRouter(*cfg)

}
