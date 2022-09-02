package main

import (
	"github.com/alekseyklimenko/go-tglist-bot/logger"
	"github.com/alekseyklimenko/go-tglist-bot/services/initservices"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//conf := config.New()
	//db := database.New(conf)
	//database.RunMigrations(db)
	initservices.Init( /*db, conf*/ )
	deferShutdown()
}

func deferShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.NewEntry().Info("Shutting down bot...")
	initservices.Shutdown()
	logger.NewEntry().Info("Bot stopped.")
}
