package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
)

func main() {
	fmt.Println("Hello, Padiwan")
	resp, err := http.Get("https://engineering-task.elancoapps.com/api/applications")
	if err != nil {
		log.Fatal(err)
	}

	// close the body after we're done with it
	defer resp.Body.Close()

	raw, _ := ioutil.ReadAll(resp.Body)

	body := string(raw)

	var arr []string

	json.Unmarshal([]byte(body), &arr)

	fmt.Printf("Found %d apps\n\n", len(arr))

	for _, app := range arr {
		fmt.Printf(" > %s\n", app)
	}
}

