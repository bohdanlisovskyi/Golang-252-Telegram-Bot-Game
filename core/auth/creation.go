package auth

import (
	"math"

	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/database"
	"github.com/bohdanlisovskyi/telegrambot/tbot"
)

func PlaceNewPlanet(amountOfPlanets int) (int, int) {
	if amountOfPlanets == 0 {
		return 0, 0
	}
	flatOfSpiral := 1.1
	cornerCoefficient := 16.0
	fAmountOfPlanets := float64(amountOfPlanets)
	//Archimedean spiral equation with some  modifications
	xCoordinate := flatOfSpiral * (fAmountOfPlanets + 1) * math.Cos(cornerCoefficient/math.Abs(math.Log10(fAmountOfPlanets+1))*math.Sqrt(fAmountOfPlanets+1))
	yCoordinate := flatOfSpiral * (fAmountOfPlanets + 1) * math.Sin(cornerCoefficient/math.Abs(math.Log10(fAmountOfPlanets+1))*math.Sqrt(fAmountOfPlanets+1))
	return int(xCoordinate), int(yCoordinate)
}

func UserCreation(message *tbot.Message) (string, error) {
	db, err := database.GetPostgresConnection()
	if err != nil {
		return "", err
	}
	plan := &database.Planet{}
	db.Last(plan)
	last_id := plan.Id
	x, y := PlaceNewPlanet(last_id)
	planet := database.Planet{PlanetName: message.Vars["planet_name"], UserName: message.From, XCoordinate: x, YCoordinate: y, IsActive: true, IsLoaded: true}
	db.Create(&planet)
	return "all ok", nil
}
