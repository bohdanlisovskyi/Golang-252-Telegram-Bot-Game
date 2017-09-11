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
		Path:        "/name {planet_name}",
		HandlerFunc: handlers.PlanetName,
	},
	Route{
		Path:        "Tutorial",
		HandlerFunc: handlers.GameTutorial,
	},
}
