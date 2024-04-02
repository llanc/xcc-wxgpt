package ability

import (
	"encoding/json"
	"fmt"
)

type Param struct {
	Description string `json:"desc"`
	Name        string `json:"name"`
}
type Ability struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"desc"`
	FormalParams []Param `json:"params"`
}
type Abilities []Ability

func (list *Abilities) ToString() (string, error) {
	jsonData, err := json.Marshal(list)
	if err != nil {
		fmt.Printf("Error marshaling to JSON: %v\n", err)
		return "", err
	}
	return string(jsonData), nil
}
