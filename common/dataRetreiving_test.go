package common

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

func TestGetResources(t *testing.T) {
	dbCrystalAmount := 111
	dbIronAmount := 222
	planetName := "myOne"
	crystalAmount, ironAmount, err := GetResources(planetName)
	if err != nil {
		t.Error(err)
	}
	if dbCrystalAmount != crystalAmount {
		t.Error("Crystal amount retreiving invalid")
	}
	if dbIronAmount != ironAmount {
		t.Error("Iron amount retreiving invalid")
	}
	t.Log("Retreivings are valid")
}

