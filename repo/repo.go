package repo

import (
	"errors"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func GetModel(dtmi string) (string, error) {
	if !IsValidDtmi(dtmi) {
		return "", errors.New("Invalid dtmi: " + dtmi)
	}
	path := DtmiToPath(dtmi)
	baseRepoUrl := "https://devicemodels.azure.com/"
	modelUrl := baseRepoUrl + path
	return get(modelUrl)
}

func DtmiToPath(dtmi string) string {
	if !IsValidDtmi(dtmi) {
		return ""
	}

	path := strings.Replace(dtmi, ":", "/", -1)
	path = strings.Replace(path, ";", "-", 1)
	path = strings.ToLower(path)
	path = path + ".json"

	return path
}

func IsValidDtmi(dtmi string) bool {
	if len(dtmi) == 0 || dtmi == "" {
		return false
	}
	matched, err := regexp.MatchString(`^dtmi:[A-Za-z](?:[A-Za-z0-9_]*[A-Za-z0-9])?(?::[A-Za-z](?:[A-Za-z0-9_]*[A-Za-z0-9])?)*;[1-9][0-9]{0,8}$`, dtmi)
	if err != nil {
		return false
	}
	return matched
}

func get(url string) (string, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	log.Println("request to:", url)
	req.Header = http.Header{
		"User-Agent": {"digimaun/odt"},
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	log.Println("response status:", resp.Status)
	log.Println("response headers:", resp.Header)
	if resp.StatusCode != 200 {
		return "", errors.New(
			"Invalid status code for model fetch: " + resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)
	return string(body), nil
}
