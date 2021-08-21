package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type people struct {
	Number int `json:"number"`
}

func main() {
	apiURL := "http://api.open-notify.org/astros.json"

	people, err := getAstros(apiURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d people found in space.\n", people.Number)
}

func getAstros(apiURL string) (people, error) {
	p := people{}

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return p, err
	}

	// Be a good net citizen:
	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return p, err
	}

	if res.Body != nil {
		// At the end of the function, close this:
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return p, err
	}

	// Unpack the JSON response into the waiting struct:
	err = json.Unmarshal(body, &p)
	if err != nil {
		return p, err
	}

	return p, nil
}
