package t3d

import (
	"log"

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
		switch v.Rune {
		case 'w':
			cameraPos.add(extend(*cameraDir, speed))
		case 'a':
			cameraPos.add(extend(zRot(*cameraDir, -90), speed))
		case 's':
			cameraPos.sub(extend(*cameraDir, speed))
		case 'd':
			cameraPos.add(extend(zRot(*cameraDir, 90), speed))
		}
	}
}
