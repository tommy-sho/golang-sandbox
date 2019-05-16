package main

import (
	"fmt"
	"github.com/tommy-sho/golang-sandbox/logger"
	"log"
)

func main() {
	// Give logging level for logger throw flag or etc.
	logLevel := "debug"
	logger, err := logger.NewLogger(logger.WithLevel(logLevel))
	if err != nil {
		log.Fatal("failed to set logger", err)
	}
	defer logger.Sync()
	fmt.Printf("%+v", logger)
}
