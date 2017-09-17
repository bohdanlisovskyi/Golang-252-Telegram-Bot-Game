package algo

import "testing"

func TestComputeIncreaseOfPeople(t *testing.T) {
	testCases := []struct {
		name                             string
		currentAmount, maxAmount, result int
	}{
		{"Zero values", 0, 0, 0},
		{"Simple case", 10, 20, 1},
		{"Max amount riched", 29, 30, 1},
		{"Not enough people to reproduce", 9, 12, 0},
		{"Current bigger then maximal", 22, 19, -3},

	}
	for _, test := range testCases {
		testLocal := test
		t.Run(testLocal.name, func(t *testing.T) {
			t.Parallel()
			increase := ComputeIncreaseOfPeople(testLocal.currentAmount, testLocal.maxAmount)
			if increase != testLocal.result {
				t.Errorf("With values %d and %d got %d , expected %d", testLocal.currentAmount, testLocal.maxAmount, increase, testLocal.result)
			}
		})
	}
}
