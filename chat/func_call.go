package chat

import (
	"encoding/json"
	"xcc-wxgpt/ability"
)

type FuncCall struct {
	Status string   `json:"status"`
	FuncId string   `json:"funcId"`
	Params []string `json:"params"`
	Guide  string   `json:"guide"`
}

func (funcCall *FuncCall) Local() {
	ability.AbilityMap[funcCall.FuncId](funcCall.Params...)
}

func GetFuncCallFromGpt(str string) (*FuncCall, error) {
	var funcCall FuncCall
	err := json.Unmarshal([]byte(str), &funcCall)
	if err != nil {
		return nil, err
	}
	return &funcCall, nil
}
