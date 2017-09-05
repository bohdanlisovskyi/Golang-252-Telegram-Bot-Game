package router

import (
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/core/handlers"
	"github.com/yanzay/tbot"
)

type Route struct {
	Type        string
	Path        string
	Reply       string
	HandlerFunc tbot.HandlerFunction
}

type Routes []Route

// build router from routs in listener.go
func NewRouter(bot *tbot.Server) *tbot.Server {

	for _, route := range routes {

		if route.Reply != "" {
			bot.Handle(route.Path, route.Reply)
			continue
		}
		bot.HandleFunc(route.Path, route.HandlerFunc)
	}

	bot.HandleDefault(handlers.EchoHandler)
	return bot
}
