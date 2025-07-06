package t3d

//TODO optimize step by removing repeating color

func Test() {

	myWorld := world{}
	ball := &sphere{}

	//BALLS========================

	//SUN
	ball = &sphere{center: vec3{x: -500, y: 0, z: -500}, radius: 300.0}
	ball.isLight = true //this makes it the sun
	ball.update = func(index int) {
		// myWorld.objects[0].center.rotateAround(vec3{x: 0, y: 0, z: 0}, 2.5, "z")
	}
	myWorld.objects = append(myWorld.objects, ball)
	myWorld.sun = myWorld.objects[0]

	//BALL
	ball = &sphere{center: vec3{x: 0, y: 30, z: 0}, radius: 5.0}
	ball.update = func(index int) {
		myWorld.objects[index].center.rotateAround(vec3{x: 0, y: 0, z: 0}, 2.5, "x")
	}
	myWorld.objects = append(myWorld.objects, ball)
	//BALL
	ball = &sphere{center: vec3{x: 0, y: 30, z: 0}, radius: 5.0}
	ball.center.rotateAround(vec3{x: 0, y: 0, z: 0}, 90, "x")
	ball.update = func(index int) {
		myWorld.objects[index].center.rotateAround(vec3{x: 0, y: 0, z: 0}, 2.5, "x")
	}
	myWorld.objects = append(myWorld.objects, ball)
	//BALL
	ball = &sphere{center: vec3{x: 0, y: 30, z: 0}, radius: 5.0}
	ball.center.rotateAround(vec3{x: 0, y: 0, z: 0}, 180, "x")
	ball.update = func(index int) {
		myWorld.objects[index].center.rotateAround(vec3{x: 0, y: 0, z: 0}, 2.5, "x")
	}
	myWorld.objects = append(myWorld.objects, ball)
	//BALL
	ball = &sphere{center: vec3{x: 0, y: 30, z: 0}, radius: 5.0}
	ball.center.rotateAround(vec3{x: 0, y: 0, z: 0}, 270, "x")
	ball.update = func(index int) {
		myWorld.objects[index].center.rotateAround(vec3{x: 0, y: 0, z: 0}, 2.5, "x")
	}
	myWorld.objects = append(myWorld.objects, ball)

	//=============== NEXT GROUP ========
	//=============== NEXT GROUP ========
	//=============== NEXT GROUP ========
	//BALL
	ball = &sphere{center: vec3{x: 100, y: 30, z: 0}, radius: 5.0}
	ball.update = func(index int) {
		myWorld.objects[index].center.rotateAround(vec3{x: 100, y: 0, z: 0}, 2.5, "x")
	}
	myWorld.objects = append(myWorld.objects, ball)
	//BALL
	ball = &sphere{center: vec3{x: 100, y: 30, z: 0}, radius: 5.0}
	ball.center.rotateAround(vec3{x: 100, y: 0, z: 0}, 90, "x")
	ball.update = func(index int) {
		myWorld.objects[index].center.rotateAround(vec3{x: 100, y: 0, z: 0}, 2.5, "x")
	}
	myWorld.objects = append(myWorld.objects, ball)
	//BALL
	ball = &sphere{center: vec3{x: 100, y: 30, z: 0}, radius: 5.0}
	ball.center.rotateAround(vec3{x: 100, y: 0, z: 0}, 180, "x")
	ball.update = func(index int) {
		myWorld.objects[index].center.rotateAround(vec3{x: 100, y: 0, z: 0}, 2.5, "x")
	}
	myWorld.objects = append(myWorld.objects, ball)
	//BALL
	ball = &sphere{center: vec3{x: 100, y: 30, z: 0}, radius: 5.0}
	ball.center.rotateAround(vec3{x: 100, y: 0, z: 0}, 270, "x")
	ball.update = func(index int) {
		myWorld.objects[index].center.rotateAround(vec3{x: 100, y: 0, z: 0}, 2.5, "x")
	}
	myWorld.objects = append(myWorld.objects, ball)

	//=============== NEXT GROUP ========
	//=============== NEXT GROUP ========
	//=============== NEXT GROUP ========
	//BALL
	ball = &sphere{center: vec3{x: 200, y: 30, z: 0}, radius: 5.0}
	ball.update = func(index int) {
		myWorld.objects[index].center.rotateAround(vec3{x: 200, y: 0, z: 0}, 2.5, "x")
	}
	myWorld.objects = append(myWorld.objects, ball)
	//BALL
	ball = &sphere{center: vec3{x: 200, y: 30, z: 0}, radius: 5.0}
	ball.center.rotateAround(vec3{x: 200, y: 0, z: 0}, 90, "x")
	ball.update = func(index int) {
		myWorld.objects[index].center.rotateAround(vec3{x: 200, y: 0, z: 0}, 2.5, "x")
	}
	myWorld.objects = append(myWorld.objects, ball)
	//BALL
	ball = &sphere{center: vec3{x: 200, y: 30, z: 0}, radius: 5.0}
	ball.center.rotateAround(vec3{x: 200, y: 0, z: 0}, 180, "x")
	ball.update = func(index int) {
		myWorld.objects[index].center.rotateAround(vec3{x: 200, y: 0, z: 0}, 2.5, "x")
	}
	myWorld.objects = append(myWorld.objects, ball)
	//BALL
	ball = &sphere{center: vec3{x: 200, y: 30, z: 0}, radius: 5.0}
	ball.center.rotateAround(vec3{x: 200, y: 0, z: 0}, 270, "x")
	ball.update = func(index int) {
		myWorld.objects[index].center.rotateAround(vec3{x: 200, y: 0, z: 0}, 2.5, "x")
	}
	myWorld.objects = append(myWorld.objects, ball)

	//=============== NEXT GROUP ========
	//=============== NEXT GROUP ========
	//=============== NEXT GROUP ========

	//BALL
	// ball = &sphere{center: vec3{x: 0.0, y: 0, z: 0}, radius: 10.0}
	// myWorld.objects = append(myWorld.objects, ball)
	// ball.update = func() {}
	//==============================

	camera := newCamera()
	myWorld.camera = camera

	myWorld.RenderFrame()
}
