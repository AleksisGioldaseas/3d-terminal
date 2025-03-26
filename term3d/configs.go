package t3d

var (
	cleanTerminal     = "\033[2J"
	moveCursorToStart = "\033[H"

	framerate = 24

	//CAMERA DEFAULTS
	cameraVerticalAngle   = 30.0
	cameraHorizontalAngle = 30.0
	cameraRayCount        = 24000
	// cameraDirection       = DIRS.Forward
	cameraDirection = DIRS.Forward
	cameraPos       = vec3{x: -20, y: 0, z: 5}
)
