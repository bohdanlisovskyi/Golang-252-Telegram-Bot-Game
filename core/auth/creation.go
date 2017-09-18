package auth

import (
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/algo"
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/database"
	"github.com/bohdanlisovskyi/telegrambot/tbot"
)

//UserCreation: creating all dependencies for new user
func UserCreation(message *tbot.Message) error {
	db, err := database.GetPostgresConnection()
	if err != nil {
		return err
	}
	plan := &database.Planet{}
	db.Last(plan)
	x, y := algo.PlaceNewPlanet(plan.Id)
	planet := database.Planet{
		PlanetName:  message.Vars["planet_name"],
		UserName:    message.From,
		XCoordinate: x,
		YCoordinate: y,
		IsActive:    true,
		IsLoaded:    false,
	}
	db.Create(&planet)
	citycenter := database.Citycenter{
		PlanetName:      message.Vars["planet_name"],
		Level:           1,
		PeopleAmount:    10,
		PeopleMaxAmount: 20,
		CitycenterKPI:   1,
	}
	db.Create(&citycenter)
	resource := database.Resource{
		PlanetName:    message.Vars["planet_name"],
		IronAmount:    500,
		CrystalAmount: 500,
	}
	db.Create(&resource)
	ironMine := database.IronMine{
		PlanetName:      message.Vars["planet_name"],
		Level:           1,
		MineKPI:         1,
		PeopleAmount:    5,
		PeopleMaxAmount: 10,
	}
	db.Create(&ironMine)
	crystalMine := database.CrystalMine{
		PlanetName:      message.Vars["planet_name"],
		Level:           1,
		MineKPI:         1,
		PeopleAmount:    5,
		PeopleMaxAmount: 10,
	}
	db.Create(&crystalMine)
	cosmodrome := database.Cosmodrome{
		PlanetName:      message.Vars["planet_name"],
		Level:           1,
		CosmodrKPI:      2,
		ShipAmount:      0,
		ShipMaxAmount:   20,
		FleetsAmount:    0,
		FleetsMaxamount: 2,
	}
	db.Create(&cosmodrome)
	return nil
}
