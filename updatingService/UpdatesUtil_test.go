package updatingService

import "testing"

func TestGetPlanetName(t *testing.T) {
	myPlanetName := "myOne"
	planetName, err := GetPlanetName("yavval")
	if err != nil {
		t.Fail()
	}
	if myPlanetName != planetName {
		t.Logf("Mistake, correct name %s, but retreived %s", myPlanetName, planetName)
		t.Fail()
	}
	if myPlanetName == planetName {
		t.Logf("Retrieving correct: %s = retreived: %s", myPlanetName, planetName)
	}
}
