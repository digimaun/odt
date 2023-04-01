package main

import (
	"fmt"

	"github.com/digimaun/odt/repo"
)

func main() {
	model := repo.GetModel("dtmi:azure:DeviceManagement:DeviceInfozzrmation;1")
	fmt.Println(model)
}
