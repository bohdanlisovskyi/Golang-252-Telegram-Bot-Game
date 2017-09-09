package algo

import (
	"math/rand"
	"time"
)

const randomRange = 0.25

type Ship struct {
	Id           int `gorm:"primary_key;AUTO_INCREMENT"`
	CosmodromeId int
	Level        int
	HitRate      int
	Health       int
	LoadCapacity int
}

//Fight - return fleets after battle and power of fleets (0 means fleet is dead, maybe they both dead)
func Fight(fleet1, fleet2 []Ship) ([]Ship, []Ship, int, int) {
	rGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		currentPowerOfFirstFleet := 0
		currentPowerOfSecondFleet := 0
		//compute summary hit of first and second fleets
		for _, ship := range fleet1 {
			if ship.Health == 0 {
				continue
			}
			if rGen.Float64() <= 0.5 {
				currentPowerOfFirstFleet += int(float64(ship.HitRate) * (1.0 + rGen.Float64()*randomRange))
			} else {
				currentPowerOfFirstFleet += int(float64(ship.HitRate) * (1.0 - rGen.Float64()*randomRange))
			}
		}
		for _, ship := range fleet2 {
			if ship.Health == 0 {
				continue
			}
			if rGen.Float64() <= 0.5 {
				currentPowerOfSecondFleet += int(float64(ship.HitRate) * (1.0 + rGen.Float64()*randomRange))
			} else {
				currentPowerOfSecondFleet += int(float64(ship.HitRate) * (1.0 - rGen.Float64()*randomRange))
			}
		}
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
	return fleet1, fleet2, 0, 0
}
