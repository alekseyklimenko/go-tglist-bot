package services

var (
	Bot BotService
)

type BotService interface {
	Shutdown()
}
