package controllers

import (
	"github.com/alekseyklimenko/go-tglist-bot/controllers/basic"
	tele "gopkg.in/telebot.v3"
)

func RegisterHandlers(b *tele.Bot) {
	basic.RegisterHandlers(b)
}
