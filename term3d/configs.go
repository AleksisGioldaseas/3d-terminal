package t3d

var (
	cleanTerminal     = "\033[2J"
	moveCursorToStart = "\033[H"

	framerate = 24

	//CAMERA DEFAULTS
	cameraVerticalAngle   = 90.0
	cameraHorizontalAngle = 90.0
	cameraRayCount        = 30000

	cameraDirection = DIRS.Forward

	cameraPos = vec3{x: -40, y: 0, z: -30}
)
