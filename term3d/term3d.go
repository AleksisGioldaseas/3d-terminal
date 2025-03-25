package t3d

//TODO optimize step by removing repeating color

func Test() {
	myWorld := world{}

	ball := &sphere{center: vec3{x: 70.0, y: 0, z: 0}, radius: 5.0}
	ball.update = func() {}

	myWorld.objects = append(myWorld.objects, ball)

	ball = &sphere{center: vec3{x: 30.0, y: 10, z: 0}, radius: 15.0}
	ball.update = func() {
		myWorld.objects[1].center.rotateAround(vec3{x: 70, y: 0, z: 0}, 1.5, "z")
	}
	myWorld.objects = append(myWorld.objects, ball)

	ball = &sphere{center: vec3{x: 82.0, y: 10, z: 0}, radius: 3.0}

	ball.update = func() {
		myWorld.objects[2].center.rotateAround(vec3{x: 70, y: 0, z: 0}, 6.0, "z")
	}
	myWorld.objects = append(myWorld.objects, ball)

	camera := newCamera()
	myWorld.camera = camera

	myWorld.sunPosition = vec3{x: 70, y: 30, z: 200}
	myWorld.srotationCenter = vec3{x: 70, y: 0, z: 0}
	myWorld.srotationSpeed = 1.0

	myWorld.RenderFrame()

	// for _, c := range pixelMap[0] {
	// 	for range 100 {
	// 		for range 100 {
	// 			fmt.Print(string(c))
	// 		}
	// 		fmt.Println()
	// 	}
	// 	time.Sleep(time.Millisecond * 700)
	// }
}
