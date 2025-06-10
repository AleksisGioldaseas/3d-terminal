package t3d

import (
	"math"
)

type vec3 struct {
	x float64
	y float64
	z float64
}

//if you don't know what matrix multiplications are look em up, I suggest 3Blue1Brown

//  x      y      z
//  1      0      0
//  0   cosθ  −sinθ
//  0   sinθ   cosθ

// rotates vector in x axis
func (vec *vec3) xRot(deg float64) {
	rads := deg2rad(deg)
	cos := math.Cos(rads)
	sin := math.Sin(rads)
	vec.y, vec.z = (cos*vec.y)+(-sin*vec.z), (sin*vec.y)+(cos*vec.z)
}

func (a vec3) cross(b vec3) vec3 {
	c := vec3{a.y*b.z - a.z*b.y,
		a.z*b.x - a.x*b.z,
		a.x*b.y - a.y*b.x}
	return c
}

func (a *vec3) perpendicular() vec3 {
	var ref vec3
	if a.x != 0 || a.z != 0 {
		ref = vec3{0, 1, 0} // world up
	} else {
		ref = vec3{1, 0, 0} // fallback if 'a' is vertical
	}
	return a.cross(ref)
}

//     x   y     z
//  cosθ   0  sinθ
//     0   1     0
// -sinθ   0  cosθ

// rotates vector in y axis
func (vec *vec3) yRot(deg float64) {
	rads := deg2rad(deg)
	cos := math.Cos(rads)
	sin := math.Sin(rads)
	vec.x, vec.z = (cos*vec.x)+(sin*vec.z), (-sin*vec.x)+(cos*vec.z)
}

//    x     y   z
// cosθ -sinθ   0
// sinθ  cosθ   0
//    0     0   1

// rotates vector in z axis
func (vec *vec3) zRot(deg float64) {
	rads := deg2rad(deg)
	cos := math.Cos(rads)
	sin := math.Sin(rads)
	vec.x, vec.y = (cos*vec.x)+(-sin*vec.y), (sin*vec.x)+(cos*vec.y)
}

// dot product
func (A *vec3) dot(B vec3) float64 {
	return A.x*B.x + A.y*B.y + A.z*B.z
}

// multiply
func (A *vec3) mult(B vec3) {
	A.x *= B.x
	A.y *= B.y
	A.z *= B.z
}

// addition
func (A *vec3) add(B vec3) {
	A.x += B.x
	A.y += B.y
	A.z += B.z
}

// subtraction
func (A *vec3) sub(B vec3) {
	A.x -= B.x
	A.y -= B.y
	A.z -= B.z
}

// normalizes the vector, basically makes the length of the vector to 1
func (vec *vec3) norm() {
	mag := vec.mag()
	vec.x /= mag
	vec.y /= mag
	vec.z /= mag
}

// magnitude, basically the length of the vector
func (vec *vec3) mag() float64 {
	return math.Sqrt(vec.x*vec.x + vec.y*vec.y + vec.z*vec.z)
}

// makes the vector longer by the scalar amount
func (vec *vec3) scale(scalar float64) {
	vec.x *= scalar
	vec.y *= scalar
	vec.z *= scalar
}

// makes the vector longer by the scalar amount
func (vec *vec3) extend(length float64) {
	mag := vec.mag()
	vec.norm()
	vec.scale(mag + length)
}

// the angle...!
func (vecA *vec3) angle(vecB vec3) float64 {
	cosAngle := vecA.dot(vecB) / (mag(*vecA) * mag(vecB))
	return cosAngle
}

// rotate the vector around a point
func (vecA *vec3) rotateAround(vecB vec3, amount float64, t string) {
	rotationVector := sub(*vecA, vecB)
	switch t {
	case "z":
		rotationVector.zRot(amount)
	case "y":
		rotationVector.yRot(amount)
	case "x":
		rotationVector.xRot(amount)
	}
	rotationVector.add(vecB)
	*vecA = rotationVector
}

var DIRS = struct {
	Up      vec3
	Down    vec3
	Forward vec3
	Back    vec3
	Right   vec3
	Left    vec3
}{
	Up:      vec3{x: 0, y: 0, z: 1},
	Down:    vec3{x: 0, y: 0, z: -1},
	Forward: vec3{x: 1, y: 0, z: 0},
	Back:    vec3{x: -1, y: 0, z: 0},
	Right:   vec3{x: 0, y: -1, z: 0},
	Left:    vec3{x: 0, y: 1, z: 0},
}
