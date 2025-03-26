package t3d

import "math"

type sphere struct {
	center         vec3
	radius         float64
	rotationCenter vec3
	rotationSpeed  float64
	update         func()
	isLight        bool
}

func (sphere *sphere) collideWithRay(raycastPosition, raycastDirection vec3) (bool, vec3) {
	dotProduct := dot(sub(sphere.center, raycastPosition), raycastDirection)

	if dotProduct <= 0 {
		return false, vec3{}
	}

	//Closest point (on raycast) to center of sphere
	ClosestPointToCenterOfSphere := add(scale(raycastDirection, dotProduct), raycastPosition)

	//distance from sphere center to CPTCOS
	S2CPTClen := mag(sub(sphere.center, ClosestPointToCenterOfSphere))

	if S2CPTClen > sphere.radius {
		return false, vec3{}
	}

	x := math.Sqrt((sphere.radius * sphere.radius) - (S2CPTClen * S2CPTClen)) //TODO investigate potential floating point impressision issues

	intersectionPoint := extend(sub(ClosestPointToCenterOfSphere, raycastPosition), -x)

	return true, intersectionPoint
}
