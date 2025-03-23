package terminal3d

import (
	"fmt"
	"math"
	"time"
)

//TODO optimize step by removing repeating color

var DIRS = struct {
	Up      pos
	Down    pos
	Forward pos
	Back    pos
	Right   pos
	Left    pos
}{
	Up:      pos{x: 0, y: 0, z: 1},
	Down:    pos{x: 0, y: 0, z: -1},
	Forward: pos{x: 1, y: 0, z: 0},
	Back:    pos{x: -1, y: 0, z: 0},
	Right:   pos{x: 0, y: -1, z: 0},
	Left:    pos{x: 0, y: 1, z: 0},
}

type pos struct {
	x float64
	y float64
	z float64
}

type camera struct {
	vAngle     float64 //vertical angle of fov
	hAngle     float64 //horizontal angle of fov
	rayCount   int     //number of rays / pixels
	rayBoxSide int     //square root of rayCount, because it represent the side of the box that each pixel is using one ray
	position   pos     //where in space the camera is located at
	direction  pos     //towards what direction the camera is looking at
	innerFrame frame   //the frame where each ray will place its value in
}

func NewCamera() camera {
	camera := camera{}
	camera.vAngle = 30
	camera.hAngle = 30
	camera.rayCount = 1000
	camera.rayBoxSide = int(math.Sqrt(float64(camera.rayCount)))
	camera.direction = DIRS.Forward

	f := frame{}
	f.canvas = make([][]pixel, camera.rayBoxSide)
	for i := range camera.rayBoxSide {
		f.canvas[i] = make([]pixel, camera.rayBoxSide)
	}
	camera.innerFrame = f
	return camera
}

type pixel struct {
	color string
	value rune
}

type frame struct {
	canvas [][]pixel
}

type sphere struct {
	center pos
	radius float64
}

type world struct {
	camera  camera
	objects []sphere
}

//if you don't know what matrix multiplications are look em up, I suggest 3Blue1Brown

//  x      y      z
//  1      0      0
//  0   cosθ  −sinθ
//  0   sinθ   cosθ

// rotates vector in x axis
func (p *pos) xRot(deg float64) {
	rads := d2r(deg)
	cos := math.Cos(rads)
	sin := math.Sin(rads)
	newY := cos*p.y + -sin*p.z
	p.z = sin*p.y + cos*p.z
	p.y = newY
}

//     x   y     z
//  cosθ   0  sinθ
//     0   1     0
// -sinθ   0  cosθ

// rotates vector in y axis
func (p *pos) yRot(deg float64) {
	rads := d2r(deg)
	cos := math.Cos(rads)
	sin := math.Sin(rads)
	newX := cos*p.x + sin*p.z
	p.z = -sin*p.x + cos*p.z
	p.x = newX
}

//    x     y   z
// cosθ -sinθ   0
// sinθ  cosθ   0
//    0     0   1

// rotates vector in z axis
func (p *pos) zRot(deg float64) {
	rads := d2r(deg)
	cos := math.Cos(rads)
	sin := math.Sin(rads)
	newX := cos*p.x + -sin*p.y
	p.y = sin*p.x + cos*p.y
	p.x = newX
}

// dot product
func (p *pos) dot(p2 pos) float64 {
	return p.x*p2.x + p.y*p2.y + p.z*p2.z
}

// multiply
func (p *pos) mult(p2 pos) {
	p.x *= p2.x
	p.y *= p2.y
	p.z *= p2.z
}

// addition
func (p *pos) add(p2 pos) {
	p.x += p2.x
	p.y += p2.y
	p.z += p2.z
}

// subtraction
func (p *pos) sub(p2 pos) {
	p.x -= p2.x
	p.y -= p2.y
	p.z -= p2.z
}

// normalizes the vector, basically makes the length of the vector to 1
func (p *pos) norm() {
	mag := p.mag()
	p.x /= mag
	p.y /= mag
	p.z /= mag
}

// magnitude, basically the length of the vector
func (p *pos) mag() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}

// makes the vector longer, meant to be used on normalized vectors
func (p *pos) scale(scalar float64) {
	p.x *= scalar
	p.y *= scalar
	p.z *= scalar
}

// turns degrees to radians
func d2r(degrees float64) float64 {
	return degrees * float64(math.Pi/180)
}

func (w *world) Fire(raycastPosition, raycastDirection pos, ball sphere) bool {
	p := ball.center
	p.sub(raycastPosition)
	dotProduct := p.dot(raycastDirection)

	if dotProduct <= 0 {
		return false
	}

	closestPointToCenter := raycastDirection
	closestPointToCenter.scale(dotProduct)
	closestPointToCenter.add(raycastPosition)

	newV := ball.center
	newV.sub(closestPointToCenter)

	mag := newV.mag()
	return mag <= ball.radius

	// if equal(mag, ball.radius) {
	// 	return true
	// }

	// return true
}

const epsilon = 1e-9

func equal(a, b float64) bool {
	return math.Abs(a-b) <= epsilon
}

func (w *world) RenderFrame() {

	hSteps := w.camera.hAngle / float64(w.camera.rayBoxSide)
	hStart := -(w.camera.vAngle / 2.0)
	hEnd := -hStart

	vSteps := w.camera.vAngle / float64(w.camera.rayBoxSide)
	vStart := -(w.camera.vAngle / 2.0)
	vEnd := -vStart
	fmt.Println(hStart, hEnd, vStart, vEnd)
	workingRaycast := pos{}
	for h := hStart; h < hEnd; h += hSteps { //h stands for horizontal rotation
		for v := vStart; v < vEnd; v += vSteps { //v stands for vertical rotation
			workingRaycast = w.camera.direction
			workingRaycast.zRot(h)
			workingRaycast.yRot(v)
			// fmt.Println(workingRaycast, h, v)
			if w.Fire(w.camera.position, workingRaycast, w.objects[0]) {
				fmt.Print("# ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

var pixelMap = []string{".:;-^~=*cirJIOd#M@"}

func Test() {
	ball := sphere{center: pos{x: 50.0, y: 0, z: 0}, radius: 10.0}
	camera := NewCamera()
	myWorld := world{}
	myWorld.objects = append(myWorld.objects, ball)
	myWorld.camera = camera

	// myWorld.RenderFrame()

	for _, c := range pixelMap[0] {
		for range 100 {
			for range 100 {
				fmt.Print(string(c))
			}
			fmt.Println()
		}
		time.Sleep(time.Millisecond * 700)
	}

}
