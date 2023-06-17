package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/hebecoding/tenant-management/infrastructure/config"
	"github.com/hebecoding/tenant-management/infrastructure/database/mongo"
)

func main() {
	// init logger
	logger := utils.NewLogger()
	defer func(logger *utils.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)

	// read in configs
	if err := config.ReadInConfig(logger); err != nil {
		logger.Fatal(err)
	}

	// init db
	db, err := mongo.NewMongoDB(
		logger, context.Background(), config.Config.DB.URL,
		"tenant-management", "tenants", "rbac",
	)
	if err != nil {
		logger.Fatal(err)
	}

	defer func(db *mongo.MongoDB) {
		err := db.Client.Disconnect(context.Background())
		if err != nil {
			logger.Fatal(err)
		}
	}(db)

	keepRunning()
}

func keepRunning() {
	// Create a channel to listen for OS signals.
	signals := make(chan os.Signal, 1)

	// Notify the channel when the application receives a SIGINT or SIGTERM signal.
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Wait for the OS signal to stop the application.
	<-signals

	fmt.Println("Stopping application...")
	// Perform any necessary cleanup or shutdown tasks here.

	// Exit the application.
	os.Exit(0)
}
