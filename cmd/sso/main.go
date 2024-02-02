package sso

import (
	"log/slog"
	"os"

	"github.com/Sombrer0Dev/auth/internal/config"
)


const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main (){
	cfg := config.MustLoad()

	_ = cfg

	// TODO init logger

  // TODO init app

	// TODO init grpc
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}



