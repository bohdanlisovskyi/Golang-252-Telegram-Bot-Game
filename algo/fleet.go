package algo

import (
	"errors"
	"math/rand"
	"time"

	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/database/postgres"
)

//FleetIsAlive - return true if exists at least one ship with Health > 0, false otherwise
func FleetIsAlive(fleet []postgres.Ship) bool {
	for index := range fleet {
		if fleet[index].Health > 0 {
			return true
		}
	}
	return false
}

//ComputeLoadCapacity - return summary load capacity of fleet
func ComputeLoadCapacity(fleet []postgres.Ship) int {
	summaryCapacity := 0
	for _, ship := range fleet {
		if ship.Health > 0 {
			summaryCapacity += ship.LoadCapacity
		}
	}
	return summaryCapacity
}

//ComputeHit - return summary hit of fleet
func ComputeHit(fleet []postgres.Ship) int {
	rGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	summaryHit := 0
	for _, ship := range fleet {
		if ship.Health == 0 {
			continue
		}
		if rGen.Float64() <= 0.5 {
			summaryHit += int(float64(ship.HitRate) * (1.0 + rGen.Float64()*randomRange))
		} else {
			summaryHit += int(float64(ship.HitRate) * (1.0 - rGen.Float64()*randomRange))
		}
	}
	return summaryHit
}

//ComputeFlightTime - return time needed to rich destination between planets
func ComputeFlightTime(fleet []postgres.Ship, startPlanet, endPlanet *postgres.Planet) (time.Duration, *error) {
	speed, ok := FindMinimalLevel(fleet)
	if !ok {
		err := errors.New("Cannot find fleet's speed")
		return 0, &err
	}
	distance := EuclDist(startPlanet.XCoordinate, startPlanet.YCoordinate, endPlanet.XCoordinate, endPlanet.YCoordinate)
	return time.Duration(distance / float64(speed)), nil
}

//FindMinimalLevel - return minimal level of ship in fleet and true, or 0, false if fleet empty
func FindMinimalLevel(fleet []postgres.Ship) (int, bool) {
	if len(fleet) == 0 {
		return 0, false
	}
	minLvl := fleet[0].Level
	for index := 1; index < len(fleet) && minLvl > 1; index++ {
		if fleet[index].Health > 0 && fleet[index].Level < minLvl {
			minLvl = fleet[index].Level
		}
	}
	return minLvl, true
}
