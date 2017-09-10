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
		Path:        "Buildings",
		HandlerFunc: handlers.ShowMenu,
	},
	Route{
		Path:        "Research Centre",
		HandlerFunc: handlers.ShowResearchMenu,
	},
	Route{
		Path:        "Mines",
		HandlerFunc: handlers.ShowMinesMenu,
	},
	Route{
		Path:        "Metal mine",
		HandlerFunc: handlers.ShowMetalMinesMenu,
	},
	Route{
		Path:        "Crystal mine",
		HandlerFunc: handlers.ShowCrystalMinesMenu,
	},
	Route{
		Path:        "City Centre",
		HandlerFunc: handlers.ShowCityCentreMenu,
	},
	Route{
		Path:        "Cosmodrome",
		HandlerFunc: handlers.ShowCosmodromeMenu,
	},
	Route{
		Path:        "Fleet",
		HandlerFunc: handlers.ShowFleetMenu,
	},
	Route{
		Path:        "Dockyard",
		HandlerFunc: handlers.ShowDockyardMenu,
	},
	Route{
		Path:        "Rating",
		HandlerFunc: handlers.ShowRatingMenu,
	},
}
