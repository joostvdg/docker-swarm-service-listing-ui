package api

import (
	"../model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Will return the service listings in the form of Docker Stacks
// Modeled by the internal model.Stacks struct
func GetStacks(url string) model.Stacks {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return model.Stacks{}
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	stacks := model.Stacks{}
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
		return model.Stacks{}
	}

	json.Unmarshal([]byte(body), &stacks)
	fmt.Printf("  > Found %d stacks\n", len(stacks))
	return stacks
}
