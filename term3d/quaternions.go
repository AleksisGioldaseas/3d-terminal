package t3d

import (
	"math"
)

type quaternion struct {
	r float64
	x float64
	y float64
	z float64
}

// angle in rads. returns quaternion that describes this rotation
func newQuaternion(angle float64, vec vec3) quaternion {
	vec.norm()
	q := quaternion{}
	sin := math.Sin(angle / 2)
	q.r = math.Cos(angle / 2)
	q.x = vec.x * sin
	q.y = vec.y * sin
	q.z = vec.z * sin
	return q
}

func (q quaternion) quaternion2Vec3() (vec3, float64) {
	vec := vec3{}

	theta := 2 * math.Acos(q.r)
	if math.Abs(q.r-1.0) < 1e-8 {
		vec.x, vec.y, vec.z = 1, 0, 0
		return vec, theta
	}

	//Use math.Sqrt(1 - q.r*q.r) instead of math.Sin(theta/2) to avoid accumulating precision errors. ???
	sinHalfTheta := math.Sin(theta / 2)

	if theta > 1e-9 {
		vec.x = q.x / sinHalfTheta
		vec.y = q.y / sinHalfTheta
		vec.z = q.z / sinHalfTheta
	}

	return vec, theta
}

func (vec *vec3) qRotate(q quaternion) {
	vecQ := quaternion{0, vec.x, vec.y, vec.z}

	newQ := q
	newQ.mult(vecQ)
	q.conjugate()
	newQ.mult(q)

	vec.x, vec.y, vec.z = newQ.x, newQ.y, newQ.z
}

func (vec *vec3) qRotUp(deg float64) {
	newQ := newQuaternion(deg2rad(deg), vec3{0, -1, 0})
	vec.qRotate(newQ)
}

func (vec *vec3) qRotLeft(deg float64) {
	newQ := newQuaternion(deg2rad(deg), vec3{0, 0, 1})
	vec.qRotate(newQ)
}

func (vec *vec3) qRotLeanRight(deg float64) {
	newQ := newQuaternion(deg2rad(deg), vec3{1, 0, 0})
	vec.qRotate(newQ)
}

// func (q *quaternion) add() {

// }

func (A *quaternion) mult(B quaternion) {
	newQ := quaternion{}
	newQ.r = A.r*B.r - A.x*B.x - A.y*B.y - A.z*B.z
	newQ.x = A.r*B.x + A.x*B.r - A.y*B.z + A.z*B.y
	newQ.y = A.r*B.y + A.x*B.z + A.y*B.r - A.z*B.x
	newQ.z = A.r*B.z - A.x*B.y + A.y*B.x + A.z*B.r
	*A = newQ
}

// func (q *quaternion) scale() { //multiply by scalar

// }

func (q *quaternion) normalize() {
	mag := math.Sqrt(q.x*q.x + q.y*q.y + q.z*q.z + q.r*q.r)
	q.x /= mag
	q.y /= mag
	q.z /= mag
	q.r /= mag
}

func (q *quaternion) conjugate() {
	q.x, q.y, q.z = -q.x, -q.y, -q.z
}
