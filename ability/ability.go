package main

import (
	"encoding/json"
	"fmt"
)

type Param struct {
	Description string `json:"desc"`
	Name        string `json:"name"`
}
type Ability struct {
	ID           string               `json:"id"`
	Name         string               `json:"name"`
	Fn           func(args ...string) `json:"-"`
	Description  string               `json:"desc"`
	FormalParams []Param              `json:"params"`
}
type List struct {
	Abilities []Ability `json:"abilities"`
}

func (list *List) ToString() (string, error) {
	jsonData, err := json.MarshalIndent(list, "", "\t")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
func (ability *Ability) ToFnSign() string {
	var params []string
	for _, param := range ability.FormalParams {
		params = append(params, param.Name)
	}
	return fmt.Sprintf("%s(%s)", ability.Name, params)
}
func (ability *Ability) ToFnDesc() string {
	return fmt.Sprintf("【Function %s provides %s with parameters %v】", ability.ToFnSign(), ability.Description, ability.FormalParams)
}
