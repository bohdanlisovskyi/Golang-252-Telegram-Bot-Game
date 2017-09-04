package router

import "github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/core/handlers"

var routes = Routes{
	Route{
		Type:  "text",
		Path:  "yo",
		Reply: "YO MEN",
	},
	Route{
		Type:        "func",
		Path:        "/start",
		HandlerFunc: handlers.StartGame,
	},
	Route{
		Type:        "func",
		Path:        "Menu",
		HandlerFunc: handlers.ShowMenu,
	},
	Route{
		Type:        "default",
		HandlerFunc: handlers.EchoHandler,
	},
}
