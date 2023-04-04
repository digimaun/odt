package main

import (
	"fmt"
	"os"

	"github.com/digimaun/odt/repo"
)

func main() {
	model, err := repo.GetModel("dtmi:com:example:TemperatureController;1")
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(model)
}
