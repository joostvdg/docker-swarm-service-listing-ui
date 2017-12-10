package model

import (
	"encoding/json"
)

// Stacks is a array of a minimal representation of a Docker Stack
// This is what we expect to retrieve from the service listing service
type Stacks []struct {
	Name     string `json:"Name"`
	Services []struct {
		Name                string `json:"Name"`
		Alias               string `json:"Alias"`
		ProxyConfigurations []struct {
			HTTPS           bool   `json:"Https"`
			MainServicePath string `json:"MainServicePath"`
			ServicePath     string `json:"ServicePath"`
			ServiceDomain   string `json:"ServiceDomain"`
			ServicePort     int    `json:"ServicePort"`
		} `json:"ProxyConfigurations"`
	} `json:"Services"`
}

// Unmarshal is a implementation of JsonStruct interface's Unmarshal
func (s *Stacks) Unmarshal(data []byte) (JsonStruct, error) {
	jsonErr := json.Unmarshal(data, &s)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return s, nil
}
