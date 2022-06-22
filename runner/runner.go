package runner

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/nuntiodev/mercury/server"
	"go.uber.org/zap"
)

func Run(logger *zap.Logger) error {
	if err := godotenv.Load(".env"); err != nil {
		logger.Warn("could not get .env")
	}
	myServer, err := server.New(context.Background(), logger)
	if err != nil {
		logger.Fatal(err.Error())
	}
	return myServer.Run()
}
