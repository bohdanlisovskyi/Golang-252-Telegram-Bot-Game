package algo

import (
	"testing"
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/database"
)

func TestFindMinimalLevel(t *testing.T) {
	testCases := []struct{
		name string
		ships []database.Ship
		expectedLvl int
		expectedOK bool
	}{
		{"Empty list of ships", []database.Ship{}, 0, false},
		{"All same lvl", []database.Ship{database.Ship{Level:2, Health:1}, database.Ship{Level:2}, database.Ship{Level:2, Health:1}}, 2, true},
		{"All different lvl", []database.Ship{database.Ship{Level:2, Health:1}, database.Ship{Level:3, Health:1}, database.Ship{Level:1, Health:1}}, 1, true},
	}
	for _, test := range testCases{
		testLocal := test
		t.Run(testLocal.name, func(t *testing.T){
			t.Parallel()
			lvl, ok := FindMinimalLevel(testLocal.ships)
			if lvl != testLocal.expectedLvl || ok != testLocal.expectedOK {
				t.Errorf("For ships %v got %d lvl, expected %d lvl", testLocal.ships, lvl, testLocal.expectedLvl)
			}
		})
	}
}