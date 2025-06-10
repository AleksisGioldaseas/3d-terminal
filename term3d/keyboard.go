package t3d

import (
	"log"
	"os"

	keyboard "github.com/eiannone/keyboard"
)

var desiredPos vec3
var desiredDir AxisFrame

func ListenKeyboard(cameraDir AxisFrame, cameraPos vec3) {
	desiredPos = cameraPos
	desiredDir = cameraDir

	c, err := keyboard.GetKeys(0)
	if err != nil {
		log.Fatal("wtf, error from keyboard package!")
	}

	speed := 10.0

	for {
		v := <-c
		// tempCamDir := *cameraDir
		switch v.Rune {
		case 'w':
			desiredPos.add(extend(desiredDir.forward, speed))
		case 'a':
			desiredPos.add(extend(desiredDir.left, speed))
		case 's':
			desiredPos.sub(extend(desiredDir.forward, speed))
		case 'd':
			desiredPos.sub(extend(desiredDir.left, speed))

		case 'i':
			desiredDir.pitch(-10)
		case 'k':
			desiredDir.pitch(10)
		case 'j':
			desiredDir.yaw(-10)
		case 'l':
			desiredDir.yaw(10)
		case 'u':
			desiredDir.roll(-10)
		case 'o':
			desiredDir.roll(10)
		case 'p':
			os.Exit(1)
		}

	}
}
