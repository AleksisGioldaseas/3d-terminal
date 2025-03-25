package t3d

var (
	cleanTerminal     = "\033[2J"
	moveCursorToStart = "\033[H"

	framerate = 24

	//CAMERA DEFAULTS
	cameraVerticalAngle   = 30.0
	cameraHorizontalAngle = 30.0
	cameraRayCount        = 20000
	cameraDirection       = DIRS.Forward
)
