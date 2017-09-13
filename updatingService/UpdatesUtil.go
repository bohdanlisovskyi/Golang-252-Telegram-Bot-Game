package updatingService

import (
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/database"
)


func GetPlanetName(userName string) (planetName string, err error)  {
	db, err := database.GetPostgresConnection()
	red, err := database.Ge
}