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

func NewRouter(bot *tbot.Server) *tbot.Server {

	for _, route := range routes {
		switch route.Type {
		case "text":
			bot.Handle(route.Path, route.Reply)
		case "func":
			bot.HandleFunc(route.Path, route.HandlerFunc)
		case "default":
			bot.HandleDefault(handlers.EchoHandler)
		}
	}
	return bot
}
