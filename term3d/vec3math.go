package t3d

import (
	"math"
)

//if you don't know what matrix multiplications are; look em up, I suggest 3Blue1Brown

//  x      y      z
//  1      0      0
//  0   cosθ  −sinθ
//  0   sinθ   cosθ

// rotates vector in x axis
func xRot(vec vec3, deg float64) vec3 {
	rads := deg2rad(deg)
	cos := math.Cos(rads)
	sin := math.Sin(rads)
	vec.y, vec.z = (cos*vec.y)+(-sin*vec.z), (sin*vec.y)+(cos*vec.z)
	return vec
}

//     x   y     z
//  cosθ   0  sinθ
//     0   1     0
// -sinθ   0  cosθ

// rotates vector in y axis
func yRot(vec vec3, deg float64) vec3 {
	rads := deg2rad(deg)
	cos := math.Cos(rads)
	sin := math.Sin(rads)
	vec.x, vec.z = (cos*vec.x)+(sin*vec.z), (-sin*vec.x)+(cos*vec.z)
	return vec
}

//    x     y   z
// cosθ -sinθ   0
// sinθ  cosθ   0
//    0     0   1

// rotates vector in z axis
func zRot(vec vec3, deg float64) vec3 {
	rads := deg2rad(deg)
	cos := math.Cos(rads)
	sin := math.Sin(rads)
	vec.x, vec.y = (cos*vec.x)+(-sin*vec.y), (sin*vec.x)+(cos*vec.y)
	return vec
}

// dot product
func dot(A, B vec3) float64 {
	return A.x*B.x + A.y*B.y + A.z*B.z
}

// multiply
func mult(A, B vec3) vec3 {
	A.x *= B.x
	A.y *= B.y
	A.z *= B.z
	return A
}

// addition
func add(A, B vec3) vec3 {
	A.x += B.x
	A.y += B.y
	A.z += B.z
	return A
}

// subtraction
func sub(A, B vec3) vec3 {
	A.x -= B.x
	A.y -= B.y
	A.z -= B.z
	return A
}

// normalizes the vector, basically makes the length of the vector to 1
func norm(vec vec3) vec3 {
	mag := vec.mag()
	vec.x /= mag
	vec.y /= mag
	vec.z /= mag
	return vec
}

// magnitude, basically the length of the vector
func mag(vec vec3) float64 {
	return math.Sqrt(vec.x*vec.x + vec.y*vec.y + vec.z*vec.z)
}

// makes the vector longer by the scalar amount
func scale(vec vec3, scalar float64) vec3 {
	vec.x *= scalar
	vec.y *= scalar
	vec.z *= scalar
	return vec
}

// makes the vector longer by the scalar amount
func extend(vec vec3, length float64) vec3 {
	mag := vec.mag()
	vec.norm()
	vec.scale(mag + length)
	return vec
}

// the angle...!
func angle(vecA, vecB vec3) float64 {
	cosAngle := vecA.dot(vecB) / (mag(vecA) * mag(vecB))
	return cosAngle
}
