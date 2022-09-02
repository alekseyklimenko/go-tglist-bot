package basic

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
)

func RegisterHandlers(b *tele.Bot) {
	b.Handle("/hello", hello)
}

func hello(c tele.Context) error {
	chatId := c.Chat().ID
	name := c.Chat().FirstName
	return c.Send(fmt.Sprintf("Hello, %s! This is basic controller. Your chat id:%d", name, chatId))
}
