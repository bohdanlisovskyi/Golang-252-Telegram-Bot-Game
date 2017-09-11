package algo

import "math"

//ManhDist - compute Manhattan distance between points (x1, y1) and (x2, y2)
func ManhDist(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}

//EuclDist - compute Euclidean distance between points (x1, y1) and (x2, y2)
func EuclDist(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2))
}
