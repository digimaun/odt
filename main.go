package main

import (
	"fmt"

	"github.com/digimaun/odt/model"
	"github.com/digimaun/odt/repo"
)

func main() {
	test := model.NewModel("dtmi:com:example:Thermostat;1")
	fmt.Println(test.Dtmi)
	fmt.Println(test.DtdlVersion)
	json, _ := test.AsJson()
	fmt.Printf("%s\n", json)

	testDtmis := []string{"dtmi:com:example:TemperatureController;1", "dtmi:com:example:TemperatureController"}
	ProcessGetModel(testDtmis)

	expectFalse := repo.IsValidDtmi("dtmi")
	fmt.Println(expectFalse)

	expectFalse = repo.IsValidDtmi("dtmi:com")
	fmt.Println(expectFalse)

	expectTrue := repo.IsValidDtmi("dtmi:com:example:abcd;1")
	fmt.Println(expectTrue)
}

func ProcessGetModel(dtmis []string) {
	for i := 0; i < len(dtmis); i++ {
		fmt.Println("Processing dtmi: " + dtmis[i])
		model, err := repo.GetModel(dtmis[i])
		if err != nil {
			fmt.Println("Error: " + err.Error())
		} else {
			fmt.Println("Success: " + model)
		}
	}
}
