package terminal3d

import (
	"fmt"
	"math"
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
	camera.vAngle = 90
	camera.hAngle = 90
	camera.rayCount = 100
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

// x  1      0      0
// y  0   cosθ  −sinθ
// z  0   sinθ   cosθ
func (p *pos) xRot(deg float64) {
	rads := d2r(deg)
	cos := math.Cos(rads)
	sin := math.Sin(rads)
	p.y = cos*p.y + -sin*p.y
	p.z = sin*p.z + cos*p.z
}

// x  cosθ   0  sinθ
// y     0   1     0
// z -sinθ   0  cosθ
func (p *pos) yRot(deg float64) {
	rads := d2r(deg)
	cos := math.Cos(rads)
	sin := math.Sin(rads)
	p.x = cos*p.x + sin*p.x
	p.z = -sin*p.z + cos*p.z
}

// x cosθ -sinθ   0
// y sinθ  cosθ   0
// z    0     0   1
func (p *pos) zRot(deg float64) {
	rads := d2r(deg)
	cos := math.Cos(rads)
	sin := math.Sin(rads)
	p.x = cos*p.x + -sin*p.x
	p.y = sin*p.y + cos*p.y
}

// turns degrees to rads
func d2r(degrees float64) float64 {
	return degrees * float64(math.Pi/180)
}

func (w *world) RenderFrame() {
	// for y := range w.camera.rayBoxSide {
	// 	for x := range w.camera.rayBoxSide {

	// 	}
	// }
}

func Test() {
	ball := sphere{center: pos{x: 50, y: 0, z: 0}, radius: 10.0}
	camera := NewCamera()
	myWorld := world{}
	myWorld.objects = append(myWorld.objects, ball)
	myWorld.camera = camera

	p := pos{x: 60, y: 40, z: 20}
	fmt.Println(p)
	p.xRot(180)
	fmt.Println(p)
	p.xRot(180)

	fmt.Println(p)
	p.yRot(180)
	fmt.Println(p)
	p.yRot(180)

	fmt.Println(p)
	p.zRot(180)
	fmt.Println(p)
	p.zRot(180)
	fmt.Println(p)

	p.xRot(180)
	p.yRot(180)
	p.zRot(180)
	fmt.Println(p)
	p.xRot(180)
	p.yRot(180)
	p.zRot(180)
	fmt.Println(p)

}
