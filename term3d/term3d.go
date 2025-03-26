package t3d

//TODO optimize step by removing repeating color

func Test() {
	myWorld := world{}
	ball := &sphere{}

	//BALLS========================

	//SUN
	ball = &sphere{center: vec3{x: 0, y: 0, z: 0}, radius: 5.0}
	ball.isLight = true //this makes it the sun
	ball.update = func() {
		myWorld.objects[0].center.rotateAround(vec3{x: 50, y: 0, z: 0}, 1.5, "z")
	}
	myWorld.objects = append(myWorld.objects, ball)
	myWorld.sun = myWorld.objects[0]
	//BALL
	ball = &sphere{center: vec3{x: 0, y: 0, z: 0}, radius: 5.0}
	ball.center.rotateAround(vec3{x: 50, y: 0, z: 0}, 120, "z")
	ball.update = func() {
		myWorld.objects[1].center.rotateAround(vec3{x: 50, y: 0, z: 0}, 1.5, "z")
	}
	myWorld.objects = append(myWorld.objects, ball)
	//BALL
	ball = &sphere{center: vec3{x: 0, y: 0, z: 0}, radius: 5.0}
	ball.center.rotateAround(vec3{x: 50, y: 0, z: 0}, 240, "z")
	ball.update = func() {
		myWorld.objects[2].center.rotateAround(vec3{x: 50, y: 0, z: 0}, 1.5, "z")
	}
	myWorld.objects = append(myWorld.objects, ball)

	//BALL
	ball = &sphere{center: vec3{x: 50.0, y: 0, z: 0}, radius: 15.0}
	ball.update = func() {
		myWorld.objects[3].center.rotateAround(vec3{x: 0, y: 0, z: 0}, 1, "z")
	}
	myWorld.objects = append(myWorld.objects, ball)
	//BALL
	ball = &sphere{center: vec3{x: 50.0, y: 0, z: 0}, radius: 15.0}
	ball.center.rotateAround(vec3{x: 0, y: 0, z: 0}, 120, "z")
	ball.update = func() {
		myWorld.objects[4].center.rotateAround(vec3{x: 0, y: 0, z: 0}, 1, "z")
	}
	myWorld.objects = append(myWorld.objects, ball)

	//BALL
	ball = &sphere{center: vec3{x: 50.0, y: 0, z: 0}, radius: 15.0}
	ball.center.rotateAround(vec3{x: 0, y: 0, z: 0}, 240, "z")
	ball.update = func() {
		myWorld.objects[5].center.rotateAround(vec3{x: 0, y: 0, z: 0}, 1, "z")
	}
	myWorld.objects = append(myWorld.objects, ball)

	//BALL
	ball = &sphere{center: vec3{x: 0.0, y: 0, z: 0}, radius: 10.0}
	myWorld.objects = append(myWorld.objects, ball)
	ball.update = func() {}
	//==============================

	camera := newCamera()
	myWorld.camera = camera

	myWorld.RenderFrame()
}
