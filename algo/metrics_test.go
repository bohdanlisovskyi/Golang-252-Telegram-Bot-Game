package algo

import (
	"testing"
	"math"
)

func TestManhDist(t *testing.T) {
	testCases := []struct{
		name string
		x1,y1,x2,y2,expectedDistance int
	}{
		{"Equals coordinates", 1, 3, 1, 3, 0},
		{"Egypt triangle", 2,3,5,-1,7},
		{"Negative coordinates", -1,-3, -5,-2,5},
	}
	for _, test := range testCases{
		testLocal := test
		t.Run(testLocal.name, func(t *testing.T){
			t.Parallel()
			distance := ManhDist(testLocal.x1, testLocal.y1, testLocal.x2, testLocal.y2)
			if distance != testLocal.expectedDistance {
				t.Errorf("For (%d;%d) and (%d;%d) got %d, expected %d", testLocal.x1, testLocal.y1, testLocal.x2, testLocal.y2, distance, testLocal.expectedDistance)
			}
		})
	}
}

func TestEuclDist(t *testing.T) {
	testCases := []struct{
		name string
		x1,y1,x2,y2 int
		expectedDistance float64
	}{
		{"Equals coordinates", 1, 3, 1, 3, 0.0},
		{"Egypt triangle", 2,3,5,-1,5.0},
		{"Negative coordinates", -1,-3, -5,-2,math.Sqrt(17.0)},
	}
	for _, test := range testCases{
		testLocal := test
		t.Run(testLocal.name, func(t *testing.T){
			t.Parallel()
			distance := EuclDist(testLocal.x1, testLocal.y1, testLocal.x2, testLocal.y2)
			if distance != testLocal.expectedDistance {
				t.Errorf("For (%d;%d) and (%d;%d) got %f , expected %f", testLocal.x1, testLocal.y1, testLocal.x2, testLocal.y2, distance, testLocal.expectedDistance)
			}
		})
	}
}