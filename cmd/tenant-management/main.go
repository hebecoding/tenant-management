package main

import (
	"context"
	"fmt"
	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/hebecoding/tenant-management/infrastructure/config"
	"github.com/hebecoding/tenant-management/infrastructure/database/mongo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// init logger
	logger := utils.NewLogger()

	// read in configs
	if err := config.ReadInConfig(logger); err != nil {
		logger.Fatal(err)
	}

	// init db
	db, err := mongo.NewMongoDB(logger, context.Background(), config.Config.DB.MongoURL,
		"tenant-management", "tenants", "rbac")
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Client.Disconnect(context.Background())
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
