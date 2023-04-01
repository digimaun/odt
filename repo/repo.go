package repo

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func GetModel(dtmi string) string {
	path := DtmiToPath(dtmi)
	baseRepoUrl := "https://devicemodels.azure.com/"
	modelUrl := baseRepoUrl + path
	return get(modelUrl)
}

func DtmiToPath(dtmi string) string {
	path := strings.Replace(dtmi, ":", "/", -1)
	path = strings.Replace(path, ";", "-", 1)
	path = strings.ToLower(path)
	path = path + ".json"

	return path
}

func get(url string) string {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	log.Println("request to:", url)
	req.Header = http.Header{
		"User-Agent": {"digimaun/odt"},
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	log.Println("response status:", resp.Status)
	log.Println("response headers:", resp.Header)
	if resp.StatusCode != 200 {
		panic(resp)
	}

	body, _ := io.ReadAll(resp.Body)
	return string(body)
}