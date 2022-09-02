package bot

import (
	"fmt"
	"github.com/alekseyklimenko/go-tglist-bot/logger"
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

type Service struct {
	bot *tele.Bot
}

func NewService( /*conf *config.Config*/ ) *Service {
	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	b.Handle("/hello", func(c tele.Context) error {
		chatId := c.Chat().ID
		name := c.Chat().FirstName
		return c.Send(fmt.Sprintf("Hello, %s! Your chat id:%d", name, chatId))
	})

	go func() {
		logger.NewEntry().Info("Starting Bot")
		b.Start()
	}()

	return &Service{
		bot: b,
	}
}

func (s *Service) Shutdown() {
	s.bot.Stop()
}
