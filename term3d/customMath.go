package t3d

import "math"

const epsilon = 1e-9

func floatEqual(a, b float64) bool {
	return math.Abs(a-b) <= epsilon
}

// turns degrees to radians
func deg2rad(degrees float64) float64 {
	return degrees * float64(math.Pi/180)
}
