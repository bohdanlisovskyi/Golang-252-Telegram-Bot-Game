package common

import "github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/database"

// GetPlanetName retrieve value planet_name from table planets
func GetPlanetName(userName string) (planetName string, err error) {
	planet := &database.Planet{}
	db, err := database.GetPostgresConnection()
	if err != nil {
		return "", err
	}
	db.Where("user_name = ?", userName).Find(planet)
	planetName = planet.PlanetName
	//fmt.Println(planetName)
	return planetName, nil
}

// GetResources retrieve current amount of iron and crystal resources for particular user
func GetResources(planetName string) (ironAmount, crystalAmount int, err error) {
	resource := &database.Resource{}
	db, err := database.GetPostgresConnection()
	if err != nil {
		return -1, -1, err
	}
	db.Where("planet_name = ?", planetName).Find(resource)
	ironAmount = resource.IronAmount
	crystalAmount = resource.CrystalAmount
	return ironAmount, crystalAmount, nil
}

// IsEnoughResources check if user has enough resources for updating
func IsEnoughResources(planetName string) (ok bool, err error) {
	return true, nil
}
