package algo

import (
	"reflect"
	"testing"

	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/database"
)

func TestPlaceNewPlanet(t *testing.T) {
	testCases := []struct {
		name                 string
		amountOfPlanets      int
		expectedX, expectedY int
	}{
		{"First planet", 0, 0, 0},
		{"Second planet", 1, 2, 0},
	}

	for _, test := range testCases {
		testLocal := test
		t.Run(testLocal.name, func(t *testing.T) {
			t.Parallel()
			x, y := PlaceNewPlanet(testLocal.amountOfPlanets)
			if x != testLocal.expectedX || y != testLocal.expectedY {
				t.Errorf("For planet %d got (%d ; %d) , expected (%d, %d)", testLocal.amountOfPlanets, x, y, testLocal.expectedX, testLocal.expectedY)
			}
		})
	}
}

func TestFindNeighbors(t *testing.T) {
	testCases := []struct {
		name              string
		mainPlanet        *database.Planet
		listOfPlanets     []database.Planet
		amountOfNeighbors int
		expectedNeighbors []database.Planet
	}{
		{"Searching for 2 planets, when exist 0", &database.Planet{XCoordinate: 0, YCoordinate: 0}, []database.Planet{}, 2, []database.Planet{}},
	}
	for _, test := range testCases {
		testLocal := test
		t.Run(testLocal.name, func(t *testing.T) {
			t.Parallel()
			listOfNeighbors := FindNeighbors(testLocal.mainPlanet, testLocal.listOfPlanets, testLocal.amountOfNeighbors)
			if !reflect.DeepEqual(listOfNeighbors, testLocal.expectedNeighbors) {
				t.Errorf("For list of %v got %v neighbors, expected %v", testLocal.listOfPlanets, listOfNeighbors, testLocal.expectedNeighbors)
			}
		})
	}

}
