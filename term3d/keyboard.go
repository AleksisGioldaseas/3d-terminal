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
		switch v.Rune {
		case 'w':
			cameraPos.add(extend(*cameraDir, speed))
		case 'a':
			cameraPos.add(mult(extend(zRot(*cameraDir, -90), speed*(1.0/cameraDir.z)), vec3{1, 1, 0}))
		case 's':
			cameraPos.sub(extend(*cameraDir, speed))
		case 'd':
			cameraPos.add(mult(extend(zRot(*cameraDir, 90), speed*(1.0/cameraDir.z)), vec3{1, 1, 0}))

		case 'i':
			cameraPos.add(mult(extend(zRot(*cameraDir, 90), speed*(1.0/cameraDir.z)), vec3{1, 1, 0}))
		case 'k':
			cameraPos.add(mult(extend(zRot(*cameraDir, 90), speed*(1.0/cameraDir.z)), vec3{1, 1, 0}))
		case 'j':
			cameraPos.add(mult(extend(zRot(*cameraDir, 90), speed*(1.0/cameraDir.z)), vec3{1, 1, 0}))
		case 'l':
			cameraPos.add(mult(extend(zRot(*cameraDir, 90), speed*(1.0/cameraDir.z)), vec3{1, 1, 0}))

		case 'p':
			os.Exit(1)
		}
		fmt.Println("pos:", *cameraPos, "\ndir:", *cameraDir, "\n\n")
	}
}
