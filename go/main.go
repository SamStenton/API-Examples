// example HTTP GET and decode
package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
)

// Given an api url, make a GET and decode it.
func fetchfromApi(apiUrl string) []byte {
	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Fatal(err)
	}

	// close the body after we're done with it - post-return
	defer resp.Body.Close()

	raw, _ := ioutil.ReadAll(resp.Body)
	return raw

}

// decode a string array in to a []string
func decodeArray(body string) []string {
	var arr []string

	json.Unmarshal([]byte(body), &arr)

	return arr
}

// loop through apps and print them nicely
func printApps(apps []string) {
	fmt.Printf("Found %d apps\n\n", len(apps))
	for _, app := range apps {
		fmt.Printf(" > %s\n", app)
	}
}

// main code init
func main() {
	apiUrl := "https://engineering-task.elancoapps.com/api/applications"

	resp := fetchfromApi(apiUrl)

	body := string(resp)

	apps := decodeArray(body)

	printApps(apps)
}

