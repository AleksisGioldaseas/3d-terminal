package t3d

import (
	"fmt"
	"strings"
	"time"
)

type world struct {
	sunPosition vec3
	camera      camera
	objects     []sphere

	srotationCenter vec3
	srotationSpeed  float64
}

type hit struct {
	distanceToCamera float64
	angle            float64
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

	closestHit := hit{distanceToCamera: 99999999999999999.9}

	for range 100000 {

		builder.WriteString(moveCursorToStart)

		for i := range w.objects {
			w.objects[i].center.rotateAround(w.objects[i].rotationCenter, w.objects[i].rotationSpeed, "z")
		}

		// w.sunPosition.rotateAround(w.srotationCenter, w.srotationSpeed, "z")
		w.sunPosition.rotateAround(w.srotationCenter, w.srotationSpeed, "y")

		// fmt.Println(object.center)
		time.Sleep(time.Millisecond * time.Duration(1000/framerate))
		for v := vStart; v < vEnd; v += vSteps { //v stands for vertical rotation
			for h := (hStart - 20); h < (hEnd + 20); h += hSteps { //h stands for horizontal rotation
				closestHit = hit{distanceToCamera: 99999999999999999.9, angle: 2.0}
				for _, object := range w.objects {
					workingRaycast = w.camera.direction
					workingRaycast.zRot(h)
					workingRaycast.yRot(v)

					//intersectionVec is from the raycast origin to the point of the intersection
					connected, intersectionVec := object.collideWithRay(w.camera.position, workingRaycast)

					if !connected {
						continue
					}

					intersectionPoint := add(w.camera.position, intersectionVec)
					sunToIntersectionPointVec := sub(intersectionPoint, w.sunPosition)
					sphereToIntersectionPoint := sub(intersectionPoint, object.center)

					angle := angle(sunToIntersectionPointVec, sphereToIntersectionPoint) + 1.0

					distanceToCamera := mag(intersectionVec)

					if closestHit.distanceToCamera > distanceToCamera {
						closestHit.distanceToCamera = distanceToCamera
						closestHit.angle = angle
					}

				}

				mult := 15.0
				builder.WriteString(string(pixelMap[min(29, max(0, int(closestHit.angle*mult)))]))
				builder.WriteString(string(pixelMap[min(29, max(0, int(closestHit.angle*mult)))]))

			}
			builder.WriteRune('\n')

		}
		fmt.Println(builder.String())
		builder.Reset()
	}

}

// var pixelMap = []string{".:;-^~=*cirJIOd#M@"}

var pixelMap = "MM###===***^^^::::...         "
