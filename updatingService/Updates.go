package updatingService

import (
	"github.com/bohdanlisovskyi/telegrambot/tbot"
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/database"
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/core/loger"
)

const (
	MAX_LEVEL                = 3
	LEVEL1_PEOPLE_MAX_AMOUNT = 10
	LEVEL2_PEOPLE_MAX_AMOUNT = 20
	LEVEL3_PEOPLE_MAX_AMOUNT = 30
	LEVEL1_KPI               = 1.1
	LEVEL2_KPI               = 1.2
	LEVEL3_KPI               = 1.2
	LEVEL2_UPDATE_TIME       = 259200 //72*60*60
	LEVEL3_UPDATE_TIME       = 172800 //48*60*60

)

func UpdateCitycenter(message *tbot.Message) (ok bool, err error) {

	citycenter := &database.Citycenter{}
	planetName, err := GetPlanetName(message.From)
	if err != nil {
		return false, err
	}
	db, err := database.GetPostgresConnection()
	if err != nil {
		return false, err
	}
	db.Where("planet_name = ?", planetName).Find(citycenter)
	//currentLevel := citycenter.Level
	//startTime := time.Now().Second()

	//switch currentLevel {
	//case 1:
	//	// Check if resources are enough:
	//	//if enough - charge resources from table resources
	//
	//	//endTime := startTime + LEVEL2_UPDATE_TIME
	//
	//	nextLevel := currentLevel + 1
	//	nextPeoplemaxamount := LEVEL2_PEOPLE_MAX_AMOUNT
	//	nextCitycenterkpi := LEVEL2_KPI
	//	time.Sleep(LEVEL2_UPDATE_TIME)
	//
	//	db.Model(&citycenter).
	//		Updates(map[string]interface{}{"level": nextLevel,
	//		"people_maxamount": nextPeoplemaxamount,
	//		"citycenter_kpi": nextCitycenterkpi})
	//	return true, nil
	//case 2:
	//	// Check if resources are enough:
	//	//if enough - charge resources from table resources
	//
	//	//endTime := startTime + LEVEL3_UPDATE_TIME
	//
	//	nextLevel := currentLevel + 1
	//	nextPeoplemaxamount := LEVEL3_PEOPLE_MAX_AMOUNT
	//	nextCitycenterkpi := LEVEL3_KPI
	//	time.Sleep(LEVEL3_UPDATE_TIME)
	//
	//	db.Model(&citycenter).
	//		Updates(map[string]interface{}{"level": nextLevel,
	//		"people_maxamount": nextPeoplemaxamount,
	//		"citycenter_kpi": nextCitycenterkpi})
	//	return true, nil
	//
	//case 3:
	//	loger.Log.Info("Maximum level has been reached")
	//	return false, nil
	//}

	return false, err

}

func UpdateIronmine(message *tbot.Message) (ok bool, err error) {

	ironMine := &database.IronMine{}
	planetName, err := GetPlanetName(message.From)
	if err != nil {
		return false, err
	}
	db, err := database.GetPostgresConnection()
	if err != nil {
		return false, err
	}
	db.Where("planet_name = ?", planetName).Find(ironMine)
	//currentLevel := ironMine.Level
	//startTime := time.Now().Second()

	//switch currentLevel {
	//case 1:
	//	// Check if resources are enough:
	//	//if enough - charge resources from table resources
	//
	//	endTime := startTime + LEVEL2_UPDATE_TIME
	//
	//case 2:
	//	// Check if resources are enough:
	//	//if enough - charge resources from table resources
	//
	//	endTime := startTime + LEVEL3_UPDATE_TIME
	//
	//case 3:
	//	loger.Log.Info("Maximum level has been reached")
	//	return false, nil
	//}
	return true, nil
}

func UpdateCrystalmine(message *tbot.Message) (ok bool, err error) {

	crystalMine := &database.CrystalMine{}
	planetName, err := GetPlanetName(message.From)
	if err != nil {
		return false, err
	}
	db, err := database.GetPostgresConnection()
	if err != nil {
		return false, err
	}

	db.Where("planet_name = ?", planetName).Find(crystalMine)
	currentLevel := crystalMine.Level
	//startTime := time.Now().Second()

	switch currentLevel {
	case 1:
		// Check if resources are enough:
		//if enough - charge resources from table resources

		//endTime := startTime + LEVEL2_UPDATE_TIME

	case 2:
		// Check if resources are enough:
		//if enough - charge resources from table resources

		//endTime := startTime + LEVEL3_UPDATE_TIME

	case 3:
		loger.Log.Info("Maximum level has been reached")
		return false, nil
	}
	return true, nil

}

func UpdateCosmodrome(message *tbot.Message) (ok bool, err error) {

	cosmodrome := &database.Cosmodrome{}
	planetName, err := GetPlanetName(message.From)
	if err != nil {
		return false, err
	}
	db, err := database.GetPostgresConnection()
	if err != nil {
		return false, err
	}

	db.Where("planet_name = ?", planetName).Find(cosmodrome)
	//currentLevel := cosmodrome.Level
	//startTime := time.Now().Second()
	//
	//switch currentLevel {
	//case 1:
	//	// Check if resources are enough:
	//	//if enough - charge resources from table resources
	//
	//	endTime := startTime + LEVEL2_UPDATE_TIME
	//
	//case 2:
	//	// Check if resources are enough:
	//	//if enough - charge resources from table resources
	//
	//	endTime := startTime + LEVEL3_UPDATE_TIME
	//
	//case 3:
	//	loger.Log.Info("Maximum level has been reached")
	//	return false, nil
	//}
	return true, nil

}

func UpdateShip(message *tbot.Message) (ok bool, err error) {

	ship := &database.Ship{}
	planetName, err := GetPlanetName(message.From)
	if err != nil {
		return false, err
	}
	db, err := database.GetPostgresConnection()
	if err != nil {
		return false, err
	}

	db.Where("planet_name = ?", planetName).Find(ship)
	//currentLevel := ship.Level
	//startTime := time.Now().Second()

	//switch currentLevel {
	//case 1:
	//	// Check if resources are enough:
	//	//if enough - charge resources from table resources
	//
	//	endTime := startTime + LEVEL2_UPDATE_TIME
	//
	//case 2:
	//	// Check if resources are enough:
	//	//if enough - charge resources from table resources
	//
	//	endTime := startTime + LEVEL3_UPDATE_TIME
	//
	//case 3:
	//	loger.Log.Info("Maximum level has been reached")
	//	return false, nil
	//}
	return true, nil

}
