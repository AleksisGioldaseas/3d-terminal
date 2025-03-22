package entry

import (
	terminal3d "3d-terminal/render"
	"log"
	"os"
)

func Start() {
	err := handleArguments()
	if err != nil {
		log.Fatal(err.Error())
	}

	terminal3d.Test()

}

func handleArguments() error {
	userArguments := os.Args[1:]
	if len(userArguments) == 0 {
		//to arguments
		return nil
	}
	return nil
}
