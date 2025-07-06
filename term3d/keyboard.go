package t3d

import (
	"log"
	"os"

	keyboard "github.com/eiannone/keyboard"
)

var vPitch = 0.0
var vYaw = 0.0
var vRoll = 0.0

var vLeft = 0.0
var vRight = 0.0
var vUp = 0.0
var vDown = 0.0

var movementApplier = func(camera *AxisFrame, pos *vec3) {
	vRoll = max(-maxValue, min(maxValue, vRoll))
	vYaw = max(-maxValue, min(maxValue, vYaw))
	vPitch = max(-maxValue, min(maxValue, vPitch))
	vUp = min(maxValue, vUp)
	vDown = min(maxValue, vDown)
	vLeft = min(maxValue, vLeft)
	vRight = min(maxValue, vRight)

	camera.roll(vRoll)
	camera.yaw(vYaw)
	camera.pitch(vPitch)

	pos.add(extend(camera.forward, vUp))
	pos.sub(extend(camera.forward, vDown))
	pos.add(extend(camera.left, vLeft))
	pos.sub(extend(camera.left, vRight))

	reduceTowardsZero(&vUp)
	reduceTowardsZero(&vLeft)
	reduceTowardsZero(&vDown)
	reduceTowardsZero(&vRight)
	reduceTowardsZero(&vPitch)
	reduceTowardsZero(&vYaw)
	reduceTowardsZero(&vRoll)

}

var decay = 0.001
var multDecay = 0.92
var posSpeed = 0.6
var rotSpeed = 0.4
var combovalue = 1.15
var maxValue = 5.0

func ListenKeyboard(cameraDir AxisFrame, cameraPos vec3) {

	c, err := keyboard.GetKeys(0)
	if err != nil {
		log.Fatal("wtf, error from keyboard package!")
	}

	for {
		v := <-c
		// tempCamDir := *cameraDir
		switch v.Rune {
		case 'w':
			vUp += posSpeed
			vUp *= combovalue
		case 'a':
			vLeft += posSpeed
			vLeft *= combovalue
		case 's':
			vDown += posSpeed
			vDown *= combovalue
		case 'd':
			vRight += posSpeed
			vRight *= combovalue

		case 'i':
			vPitch -= rotSpeed
			vPitch *= combovalue
		case 'k':
			vPitch += rotSpeed
			vPitch *= combovalue
		case 'j':
			vYaw -= rotSpeed
			vYaw *= combovalue
		case 'l':
			vYaw += rotSpeed
			vYaw *= combovalue
		case 'u':
			vRoll -= rotSpeed
			vRoll *= combovalue
		case 'o':
			vRoll += rotSpeed
			vRoll *= combovalue
		case 'p':
			os.Exit(1)
		case 't':
			multFish += 0.02
		case 'g':
			multFish -= 0.02
		}
	}
}

var reduceTowardsZero = func(v *float64) {
	if *v > 0 {
		*v = *v * multDecay
		*v -= decay
		if *v < 0 {
			*v = 0
		}
	} else if *v < 0 {
		*v = *v * multDecay
		*v += decay
		if *v > 0 {
			*v = 0
		}
	}
}
