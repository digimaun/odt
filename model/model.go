package model

import (
	"encoding/json"
	"fmt"

	"github.com/digimaun/odt/repo"
)

type Model struct {
	Dtmi        string
	Raw         string
	DtdlVersion int
	DisplayName string
	Description string
	Telemetry   map[string]map[string]string
}

type jsonPayload struct {
	Dtmi        string                       `json:"dtmi"`
	DtdlVersion int                          `json:"dtdlVersion"`
	DisplayName string                       `json:"displayName"`
	Description string                       `json:"description"`
	Telemetry   map[string]map[string]string `json:"telemetry"`
}

func (m *Model) AsJson() ([]byte, error) {
	jp := new(jsonPayload)
	jp.Description = m.Description
	jp.DisplayName = m.DisplayName
	jp.Telemetry = m.Telemetry
	jp.Dtmi = m.Dtmi
	jp.DtdlVersion = m.DtdlVersion

	return json.MarshalIndent(jp, "", "  ")
}

func NewModel(dtmi string) *Model {
	raw, err := repo.GetModel(dtmi)
	if err != nil {
		panic(err)
	}
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(raw), &dat); err != nil {
		panic(err)
	}

	m := new(Model)
	m.Dtmi = dtmi
	m.DtdlVersion = 2
	m.Raw = raw

	if val, exists := dat["displayName"]; exists {
		m.DisplayName = val.(string)
	}
	if val, exists := dat["description"]; exists {
		m.Description = val.(string)
	}

	m.Telemetry = fetchTelemetry(dat["contents"].([]interface{}))
	return m
}

func fetchTelemetry(contents []interface{}) map[string]map[string]string {
	//fmt.Println(contents)
	r := make(map[string]map[string]string)
	kpis := []string{"displayName", "description", "schema", "unit"}
	for i := 0; i < len(contents); i++ {
		attr := contents[i].(map[string]interface{})
		var name string
		process := false
		switch t := attr["@type"].(type) {
		case string:
			if attr["@type"] == "Telemetry" {
				process = true
				name = attr["name"].(string)
			}
		case []interface{}:
			if attr["@type"].([]interface{})[0] == "Telemetry" {
				process = true
				name = attr["name"].(string)
			}
		default:
			fmt.Printf("unexpected type %T\n", t)

		}
		if process {
			r[name] = make(map[string]string)
			for j := 0; j < len(kpis); j++ {
				if val, exists := attr[kpis[j]]; exists {
					r[name][kpis[j]] = val.(string)
				}
			}
		}
	}
	return r
}
