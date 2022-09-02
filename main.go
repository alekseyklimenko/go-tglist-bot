package main

import (
	"fmt"
	"github.com/alekseyklimenko/go-tglist-bot/logger"
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//conf := config.New()
	//db := database.New(conf)
	//database.RunMigrations(db)
	//initservices.Init(db, conf)

	//todo make bot as one of services
	bot := startBot( /*conf*/ )
	deferShutdown(bot)
}

func startBot( /*conf *config.Config*/ ) *tele.Bot {
	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	bot.Handle("/hello", func(c tele.Context) error {
		chatId := c.Chat().ID
		name := c.Chat().FirstName
		return c.Send(fmt.Sprintf("Hello, %s! Your chat id=%d", name, chatId))
	})

	go func() {
		logger.NewEntry().Info("Starting Bot")
		bot.Start()
	}()
	return bot
}

func deferShutdown(bot *tele.Bot) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.NewEntry().Info("Shutting down bot...")
	//	initservices.Shutdown()
	bot.Stop()
	logger.NewEntry().Info("Bot stopped.")
}
