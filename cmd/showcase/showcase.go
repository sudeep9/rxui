package main

import (
	"log/slog"
	"os"

	"github.com/sudeep9/webserver"
)

func main() {
	handlers := &Handlers{}

	opts := webserver.ServerOptions{
		Handlers: map[string]webserver.Handlers{
			"/": handlers,
		},
		StaticDirs: map[string]string{
			"/static": "./static",
		},
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	srv := webserver.NewServer(logger, opts)
	if err := srv.Start(); err != nil {
		logger.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}
