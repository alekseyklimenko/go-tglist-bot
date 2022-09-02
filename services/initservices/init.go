package initservices

import (
	"github.com/alekseyklimenko/go-tglist-bot/services"
	"github.com/alekseyklimenko/go-tglist-bot/services/bot"
)

func Init( /* db *gorm.DB, conf *config.Config */ ) {
	services.Bot = bot.NewService()
}

func Shutdown() {
	services.Bot.Shutdown()
}
