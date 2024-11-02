package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Vasudev-2308/gostudy/intenal/config"
	"github.com/Vasudev-2308/gostudy/intenal/http/handlers/Teacher"
	Student "github.com/Vasudev-2308/gostudy/intenal/http/handlers/Student"
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
	router := http.NewServeMux()
	router.HandleFunc("GET /api/students", Student.GetStudents())
	router.HandleFunc("POST /api/create", Student.CreateStudent())
	router.HandleFunc("GET /api/teachers", Teacher.GetTeacher())
	router.HandleFunc("POST /api/create-teachers", Teacher.CreateTeacher())
	
	slog.Info("Server Listening on %s 🔥", slog.String("Address: ", cfg.HttpServer.Addr))
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

	// Setup Database

}
