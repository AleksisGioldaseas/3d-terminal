package entry

import (
	t3d "3d-terminal/term3d"
	"log"
	"os"
)

func Start() {
	err := handleArguments()
	if err != nil {
		log.Fatal(err.Error())
	}

	t3d.Test()

}

func handleArguments() error {
	userArguments := os.Args[1:]
	if len(userArguments) == 0 {
		//to arguments
		return nil
	}
	return nil
}
