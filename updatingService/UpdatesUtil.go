package updatingService

import (
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/database"

)

// GetPlanetName retrieve value planet_name from table planets
func GetPlanetName(userName string) (planetName string, err error)  {
	planet := &database.Planet{}
	db, err := database.GetPostgresConnection()
	if err != nil{
		return "", err
	}
	db.Where("user_name = ?", userName).Find(planet)
	planetName = planet.PlanetName
	//fmt.Println(planetName)
	return planetName, nil
}


func GetResources(planetName string){

}


// IsEnoughResources check if user has enough resources for updating
func IsEnoughResources(planetName string) (ok bool, err error) {
	
}