package auth

import (
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/algo"
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/database"
	"github.com/bohdanlisovskyi/telegrambot/tbot"
)

func UserCreation(message *tbot.Message) error {
	db, err := database.GetPostgresConnection()
	if err != nil {
		return err
	}
	plan := &database.Planet{}
	db.Last(plan)
	last_id := plan.Id
	x, y := algo.PlaceNewPlanet(last_id)
	planet := database.Planet{PlanetName: message.Vars["planet_name"], UserName: message.From, XCoordinate: x, YCoordinate: y, IsActive: true, IsLoaded: true}
	db.Create(&planet)
	citycenter := database.Citycenter{
		PlanetName:      message.Vars["planet_name"],
		Level:           1,
		PeopleAmount:    10,
		PeopleMaxAmount: 100,
		CitycenterKPI:   2,
	}
	db.Create(&citycenter)
	resource := database.Resource{
		PlanetName:    message.Vars["planet_name"],
		IronAmount:    500,
		CrystalAmount: 500,
	}
	db.Create(&resource)
	iron_mine := database.IronMine{
		PlanetName:      message.Vars["planet_name"],
		Level:           0,
		MineKPI:         2,
		PeopleAmount:    0,
		PeopleMaxAmount: 0,
	}
	db.Create(&iron_mine)
	crystal_mine := database.CrystalMine{
		PlanetName:      message.Vars["planet_name"],
		Level:           0,
		MineKPI:         2,
		PeopleAmount:    0,
		PeopleMaxAmount: 0,
	}
	db.Create(&crystal_mine)
	cosmodrome := database.Cosmodrome{
		PlanetName:      message.Vars["planet_name"],
		Level:           0,
		CosmodrKPI:      2,
		ShipAmount:      0,
		ShipMaxAmount:   0,
		FleetsAmount:    0,
		FleetsMaxamount: 0,
	}
	db.Create(&cosmodrome)
	return nil
}
