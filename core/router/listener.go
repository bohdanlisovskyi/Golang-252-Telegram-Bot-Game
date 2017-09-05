package router

import "github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/core/handlers"

var routes = Routes{
	Route{
		Path:  "yo",
		Reply: "/start",
	},
	Route{
		Path:        "/start",
		HandlerFunc: handlers.StartGame,
	},
	Route{
		Path:        "Menu",
		HandlerFunc: handlers.ShowMenu,
	},
}
