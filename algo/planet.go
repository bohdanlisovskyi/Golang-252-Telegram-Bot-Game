package algo

import (
	"math"
)

type Planet struct {
	Id          int    `gorm:"primary_key;AUTO_INCREMENT"`
	PlanetName  string `gorm:"not null;unique"`
	XCoordinate int
	YCoordinate int
	IsActive    bool
	IsLoaded    bool
}

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

func FindNeighbors(mainPlanet *Planet, listOfPlanets []Planet, amount int) []Planet {
	var listOfNeighbors = []Planet{}
	minRadius := 1
	maxRadius := 4
	increase := maxRadius - minRadius
	for len(listOfNeighbors) < amount && len(listOfNeighbors) < len(listOfPlanets)-1 {
		for _, planet := range listOfPlanets {
			dist := ManhDist(planet.XCoordinate, planet.YCoordinate, mainPlanet.XCoordinate, mainPlanet.YCoordinate)
			if dist >= minRadius && dist < maxRadius {
				listOfNeighbors = append(listOfNeighbors, planet)
				if len(listOfNeighbors) == amount {
					return listOfNeighbors
				}
			}
		}
		minRadius += increase
		maxRadius += increase
	}
	return listOfNeighbors
}
