package t3d

var (
	cleanTerminal     = "\033[2J"
	moveCursorToStart = "\033[H"

	framerate = 24

	//CAMERA DEFAULTS
	cameraVerticalAngle   = 90.0
	cameraHorizontalAngle = 90.0
	cameraRayCount        = 30000

	cameraAxisFrame = AxisFrame{DIRS.Forward, DIRS.Up, DIRS.Left}

	cameraPos = vec3{x: -100, y: 0, z: 0}
)
