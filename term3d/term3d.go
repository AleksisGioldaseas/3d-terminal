package t3d

//TODO optimize step by removing repeating color

func Test() {
	ball := sphere{center: vec3{x: 50.0, y: 0, z: 0}, radius: 10.0}
	camera := newCamera()
	myWorld := world{}
	myWorld.objects = append(myWorld.objects, ball)
	myWorld.camera = camera

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
