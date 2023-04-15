package main

import (
	"fmt"

	"github.com/digimaun/odt/model"
	"github.com/digimaun/odt/repo"
)

func main() {
	test1 := model.NewModel("dtmi:com:example:temperaturecontroller;1")
	fmt.Println(test1.Dtmi)
	fmt.Println(test1.DtdlVersion)
	json1, _ := test1.AsJson()
	fmt.Printf("%s\n", json1)

	test2 := model.NewModel("dtmi:com:example:azuresphere:altair;1")
	fmt.Println(test2.Dtmi)
	fmt.Println(test2.DtdlVersion)
	json2, _ := test2.AsJson()
	fmt.Printf("%s\n", json2)

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
