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
	Route{
		Path:        "DataCenter",
		HandlerFunc: handlers.DataCenter,
	},
	Route{
		Path:        "Skip Tutorial",
		HandlerFunc: handlers.SkipTutorial,
	},
	Route{
		Path:        "Back to Buildings",
		HandlerFunc: handlers.AllBuildings,
	},
	Route{
		Path:        "Status",
		HandlerFunc: handlers.DataCenterStatus,
	},
	Route{
		Path:        "Help/Toturial",
		HandlerFunc: handlers.DataCenterToturial,
	},
	Route{
		Path:        "Search planets",
		HandlerFunc: handlers.DataCenterSearchPlanet,
	},
	Route{
		Path:        "List of planets",
		HandlerFunc: handlers.DataCenterSearchListOfPlanet,
	},
	Route{
		Path:        "Back to DataCenter",
		HandlerFunc: handlers.DataCenter,
	},
	Route{
		Path:        "Cosmodrom",
		HandlerFunc: handlers.CosmodromMenu,
	},
}
