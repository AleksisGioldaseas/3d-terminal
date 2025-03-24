package t3d

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type world struct {
	sunPosition vec3
	camera      camera
	objects     []sphere
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

	var builder strings.Builder
	builder.WriteString(cleanTerminal)
	workingRaycast := vec3{}
	for range 100000 {
		builder.WriteString(moveCursorToStart)

		w.objects[0].center.rotateAround(vec3{x: 70, y: 0, z: 0}, 5.0, "z")
		w.objects[0].center.rotateAround(vec3{x: 70, y: 0, z: 0}, 1.21, "x")

		fmt.Println(w.objects[0].center)
		time.Sleep(time.Millisecond * time.Duration(1000/framerate))
		for v := vStart; v < vEnd; v += vSteps { //v stands for vertical rotation
			for h := (hStart - 20); h < (hEnd + 20); h += hSteps { //h stands for horizontal rotation
				workingRaycast = w.camera.direction
				workingRaycast.zRot(h)
				workingRaycast.yRot(v)

				//intersectionVec is from the raycast origin to the point of the intersection
				connected, intersectionVec := w.fire(w.camera.position, workingRaycast, w.objects[0])

				if !connected {
					builder.WriteString("  ")
					continue
				}

				intersectionPoint := add(w.camera.position, intersectionVec)
				sunToIntersectionPointVec := sub(intersectionPoint, w.sunPosition)
				sphereToIntersectionPoint := sub(intersectionPoint, w.objects[0].center)

				ang := angle(sunToIntersectionPointVec, sphereToIntersectionPoint) + 1.0
				mult := 15.0
				builder.WriteString(string(pixelMap[min(29, max(0, int(ang*mult)))]))
				builder.WriteString(string(pixelMap[min(29, max(0, int(ang*mult)))]))

			}
			builder.WriteRune('\n')
		}
		fmt.Println(builder.String())
		builder.Reset()
	}

}

// var pixelMap = []string{".:;-^~=*cirJIOd#M@"}

var pixelMap = "MM###===***^^^::::...         "
