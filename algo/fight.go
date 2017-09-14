package algo

import (
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/database"
)

const randomRange = 0.25

//Fight - return fleets after battle and power of fleets (0 means fleet is dead, maybe they both dead)
func Fight(fleet1, fleet2 []database.Ship) ([]database.Ship, []database.Ship, int, int) {

	for {
		//compute summary hit of first and second fleets
		currentPowerOfFirstFleet := ComputeHit(fleet1)
		currentPowerOfSecondFleet := ComputeHit(fleet2)

		if currentPowerOfFirstFleet == 0 || currentPowerOfSecondFleet == 0 {
			return fleet1, fleet2, currentPowerOfFirstFleet, currentPowerOfSecondFleet
		}
		//apply hits to ships
		for index := range fleet2 {
			currentPowerOfFirstFleet -= fleet2[index].Health
			//power < 0 mean power of fleet ended and ship is alive
			if currentPowerOfFirstFleet < 0 {
				fleet2[index].Health = -currentPowerOfFirstFleet
				break
			}
			fleet2[index].Health = 0
		}
		for index := range fleet1 {
			currentPowerOfSecondFleet -= fleet1[index].Health
			if currentPowerOfSecondFleet < 0 {
				fleet1[index].Health = -currentPowerOfSecondFleet
				break
			}
			fleet1[index].Health = 0
		}
	}
}
