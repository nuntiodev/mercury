package main

import (
	"github.com/nuntiodev/mercury/runner"
	"go.uber.org/zap"
	"log"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err.Error())
	}
	runner.Run(logger)
}
