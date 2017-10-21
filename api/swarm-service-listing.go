package api

import (
	"fmt"
	"encoding/json"
	"net/http"
	"../model"
	"io/ioutil"
	"log"
)

func GetStacks(url string) model.Stacks{
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
	return stacks
}


