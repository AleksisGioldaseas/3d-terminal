package t3d

type pixel struct {
	color string
	value rune
}

type frame struct {
	canvas [][]pixel
}

type rotator struct {
	center vec3
	speed  float64
}
