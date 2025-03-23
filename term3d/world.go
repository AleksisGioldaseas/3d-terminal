package t3d

import (
	"fmt"
	"math"
)

type world struct {
	camera  camera
	objects []sphere
}

func (w *world) fire(raycastPosition, raycastDirection vec3, ball sphere) (bool, vec3) {

	dotProduct := dot(sub(ball.center, raycastPosition), raycastDirection)

	if dotProduct <= 0 {
		return false, vec3{}
	}

	//Closest point (on raycast) to center of sphere
	ClosestPointToCenterOfSphere := add(scale(raycastDirection, dotProduct), raycastPosition)

	//distance from sphere center to CPTCOS
	S2CPTClen := mag(sub(ball.center, ClosestPointToCenterOfSphere))

	if S2CPTClen > ball.radius {
		return false, vec3{}
	}

	x := math.Sqrt((ball.radius * ball.radius) - (S2CPTClen * S2CPTClen)) //TODO investigate potential floating point impressision issues

	intersectionPoint := extend(ClosestPointToCenterOfSphere, -x)

	return true, intersectionPoint
}

func (w *world) RenderFrame() {

	hSteps := w.camera.hAngle / float64(w.camera.rayBoxSide)
	hStart := -(w.camera.vAngle / 2.0)
	hEnd := -hStart

	vSteps := w.camera.vAngle / float64(w.camera.rayBoxSide)
	vStart := -(w.camera.vAngle / 2.0)
	vEnd := -vStart
	fmt.Println(hStart, hEnd, vStart, vEnd)

	workingRaycast := vec3{}
	for v := vStart; v < vEnd; v += vSteps { //v stands for vertical rotation
		for h := hStart; h < hEnd; h += hSteps { //h stands for horizontal rotation

			workingRaycast = w.camera.direction
			workingRaycast.zRot(h)
			workingRaycast.yRot(v)
			// fmt.Println(workingRaycast, h, v)
			// connected, intersectionVec := w.fire(w.camera.position, workingRaycast, w.objects[0])
			connected, _ := w.fire(w.camera.position, workingRaycast, w.objects[0])
			if connected {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

var pixelMap = []string{".:;-^~=*cirJIOd#M@"}
