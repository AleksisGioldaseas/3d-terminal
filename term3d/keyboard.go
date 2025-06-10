package t3d

import (
	"fmt"
	"log"
	"os"

	keyboard "github.com/eiannone/keyboard"
)

func ListenKeyboard(cameraDir *vec3, cameraPos *vec3) {
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
			cameraPos.add(extend(*cameraDir, speed))
		case 'a':
			cameraPos.add(extend(qRotate(*cameraDir, newQuaternion(-90, perpendicular(*cameraDir))), speed))
		case 's':
			cameraPos.sub(extend(*cameraDir, speed))
		case 'd':
			cameraPos.add(extend(qRotate(*cameraDir, newQuaternion(90, perpendicular(*cameraDir))), speed))

		case 'i':
			cameraDir.yRot(10)
		case 'k':
			cameraDir.yRot(-10)
		case 'j':
			cameraDir.zRot(10)
		case 'l':
			cameraDir.zRot(-10)

		case 'p':
			os.Exit(1)
		}
		fmt.Println("pos:", *cameraPos, "\ndir:", *cameraDir, "\n\n")
	}
}
