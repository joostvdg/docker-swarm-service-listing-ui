package model

import (
	"encoding/json"
)

type Stacks []struct {
	Name     string `json:"Name"`
	Services []struct {
		Name                string `json:"Name"`
		Alias               string `json:"Alias"`
		ProxyConfigurations []struct {
			HTTPS         bool   `json:"Https"`
			ServicePath   string `json:"ServicePath"`
			ServiceDomain string `json:"ServiceDomain"`
			ServicePort   int    `json:"ServicePort"`
		} `json:"ProxyConfigurations"`
	} `json:"Services"`
}
func (s *Stacks) Unmarshal(data []byte) (JsonStruct, error) {
	jsonErr := json.Unmarshal(data, &s)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return s, nil
}