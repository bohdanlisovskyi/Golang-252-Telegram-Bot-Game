package updatingService

import "testing"

func TestGetPlanetName(t *testing.T) {

	myPlanetName := "nljlkmk"

	planetName, err := GetPlanetName("yavval")

	if err != nil{
		//t.Fail()
	}
	if myPlanetName == planetName{
		t.Logf("Retreiving succesed, %s, %s", myPlanetName, planetName)
	}

}

