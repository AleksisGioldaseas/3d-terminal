package t3d

type pixel struct {
	color string
	value rune
}

type frame struct {
	canvas [][]pixel
}

type sphere struct {
	center vec3
	radius float64
}

type rotator struct {
	center vec3
	speed  float64
}
