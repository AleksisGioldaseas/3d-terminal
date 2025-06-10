package t3d

import (
	"fmt"
	"strings"
	"time"
)

type world struct {
	camera  camera
	objects []*sphere
	sun     *sphere
}

// type hit struct {
// 	distanceToCamera float64
// 	angle            float64
// }

func (w *world) RenderFrame() {

	go ListenKeyboard(w.camera.direction, w.camera.position)

	mult := 0.7 //multiplier to stop fish eye lense

	hSteps := w.camera.hAngle / float64(w.camera.rayBoxSide)
	hStart := (-(w.camera.vAngle))
	hEnd := (-hStart)

	vSteps := w.camera.vAngle / float64(w.camera.rayBoxSide)
	vStart := -(w.camera.vAngle / 2.0)
	vEnd := (-vStart)

	var builder strings.Builder
	builder.WriteString(cleanTerminal)

	for range 100000 {
		builder.WriteString(moveCursorToStart)

		// a + (b — a) * t

		w.camera.position = lerpVec3(w.camera.position, desiredPos, 0.2)
		w.camera.direction.forward = lerpVec3(w.camera.direction.forward, desiredDir.forward, 0.2)
		w.camera.direction.up = lerpVec3(w.camera.direction.up, desiredDir.up, 0.2)
		w.camera.direction.left = lerpVec3(w.camera.direction.left, desiredDir.left, 0.2)
		// w.camera.position.z += 1
		// w.camera.direction.rotateAround(vec3{}, 1, "y")

		for i := range w.objects {
			w.objects[i].update()
			// w.objects[i].center.rotateAround(w.objects[i].rotationCenter, w.objects[i].rotationSpeed, "z")
		}

		// w.sun.center.rotateAround(w.sun.rotationCenter, w.sun.rotationSpeed, "z")
		// w.sunPosition.rotateAround(w.srotationCenter, w.srotationSpeed, "y")

		fmt.Println(w.camera.direction)

		time.Sleep(time.Millisecond * time.Duration(1000/framerate))

		for v := vStart; v < vEnd; v += vSteps { //v stands for vertical rotation
			for h := (hStart); h < (hEnd); h += hSteps { //h stands for horizontal rotation
				//debug, move camera as you wish

				workingDirection := AxisFrame{w.camera.direction.forward, w.camera.direction.up, w.camera.direction.left}
				workingDirection.pitch(v * mult)
				workingDirection.yaw(h * mult)

				//checking collision of camera ray to first object
				collisionPoint, normalVec, collided, sphereRef := collideRayToObjects(w.camera.position, workingDirection.forward, false, w.objects)

				if !collided {
					builder.WriteString("  ")
					continue
				}

				if sphereRef.isLight {
					builder.WriteString("MM")
					continue
				}

				//checking collision towards sun
				_, _, collided, sphereRef = collideRayToObjects(collisionPoint, norm(sub(w.sun.center, collisionPoint)), false, w.objects)

				if collided && !sphereRef.isLight {
					// fmt.Println("YO")
					// time.Sleep(time.Millisecond * 50)
					builder.WriteString("  ")
					continue
				}

				angle := angle(sub(collisionPoint, w.sun.center), normalVec)

				builder.WriteString(angleToPixel(angle))

			}

			builder.WriteRune('\n')

		}

		fmt.Println(builder.String())
		builder.Reset()
	}

}

// var pixelMap = []string{".:;-^~=*cirJIOd#M@"}

// var pixelMap = "MM###===***^^^::::...         "
// var pixelMap = "MM###===***^^^                "
var pixelMap = "#aa===--::....                "

// func isPointOccluded(point, sun vec3) bool {
// 	return false
// }

func angleToPixel(angle float64) string {
	// fmt.Println(angle)
	val := min(29, max(0, int((angle+1.0)*15.0))) //TODO make this dynamically pick values from available characters
	return string(pixelMap[val]) + string(pixelMap[val])
}

func collideRayToObjects(rayOrigin, rayDirection vec3, stopOnFirstCollision bool, objects []*sphere) (vec3, vec3, bool, *sphere) {
	closestHit := 99999999999999999.9

	collided := false
	var normalVec vec3
	var intersectionPoint vec3
	var distanceToCamera float64
	var collidedSphere *sphere

	for _, object := range objects {
		//intersectionVec is from the raycast origin to the point of the intersection
		connected, intersectionVec := object.collideWithRay(rayOrigin, rayDirection)

		if !connected {
			continue
		}

		distanceToCamera = mag(intersectionVec)

		if closestHit > distanceToCamera {
			intersectionPoint = add(rayOrigin, intersectionVec)
			normalVec = norm(sub(intersectionPoint, object.center)) //the normal vector to calculate angle with the sun afterwards
			closestHit = distanceToCamera
			collided = true
			collidedSphere = object
			if stopOnFirstCollision {
				return intersectionPoint, normalVec, true, collidedSphere
			}
			//TODO if objects are sorted correctly we can probably escape early!!!
			// return intersectionPoint, normalVec, collided
		}
	}

	return intersectionPoint, normalVec, collided, collidedSphere
}
