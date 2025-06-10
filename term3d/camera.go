package t3d

import "math"

type camera struct {
	vAngle     float64   //vertical angle of fov
	hAngle     float64   //horizontal angle of fov
	rayCount   int       //number of rays / pixels
	rayBoxSide int       //square root of rayCount, because it represent the side of the box that each pixel is using one ray
	position   vec3      //where in space the camera is located at
	direction  AxisFrame //towards what direction the camera is looking at
	innerFrame frame     //the frame where each ray will place its value in
}

type AxisFrame struct {
	forward vec3 //towards what direction the camera is looking at
	up      vec3
	left    vec3
}

func (f *AxisFrame) pitch(degs float64) {
	rads := deg2rad(degs)
	f.forward.qRotate(newQuaternion(rads, f.left))
	f.up.qRotate(newQuaternion(rads, f.left))
}

func (f *AxisFrame) yaw(degs float64) {
	rads := deg2rad(degs)
	f.forward.qRotate(newQuaternion(rads, f.up))
	f.left.qRotate(newQuaternion(rads, f.up))
}

func (f *AxisFrame) roll(degs float64) {
	rads := deg2rad(degs)
	f.left.qRotate(newQuaternion(rads, f.forward))
	f.up.qRotate(newQuaternion(rads, f.forward))
}

func newCamera() camera {
	camera := camera{}
	camera.vAngle = cameraVerticalAngle
	camera.hAngle = cameraHorizontalAngle
	camera.rayCount = cameraRayCount
	camera.rayBoxSide = int(math.Sqrt(float64(camera.rayCount)))
	camera.direction = cameraAxisFrame

	camera.position = cameraPos

	// camera.direction.pitch(-30)

	f := frame{}
	f.canvas = make([][]pixel, camera.rayBoxSide)
	for i := range camera.rayBoxSide {
		f.canvas[i] = make([]pixel, camera.rayBoxSide)
	}
	camera.innerFrame = f
	return camera
}
