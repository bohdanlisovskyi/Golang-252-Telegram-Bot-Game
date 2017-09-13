package updatingService

import (
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/database"
	"fmt"
)


func GetPlanetName(userName string) (planetName string, err error)  {
	planet := &database.Planet{}
	db, err := database.GetPostgresConnection()
	if err != nil{
		return "", err
	}
	db.Where("user_name = ?", userName).Find(planet)
	planetName = planet.PlanetName
	fmt.Println(planetName)
	return planetName, nil
}