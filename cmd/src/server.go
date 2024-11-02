package router

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func StartServer(server *http.Server, doneChannel chan os.Signal) {
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
