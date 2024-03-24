package ability

import (
	"fmt"
	"strings"
)

type Param struct {
	Description string
	Name        string
}
type Ability struct {
	Name         string
	Fn           func(args ...string)
	Description  string
	FormalParams []Param
}
type List struct {
	Abilities []Ability
}

func (list *List) ToString() string {
	var checklist strings.Builder
	for _, ability := range list.Abilities {
		checklist.WriteString(ability.ToFnDesc())
		checklist.WriteString("\n")
	}
	result, _ := strings.CutSuffix(checklist.String(), "\n")
	return result
}

func (ability *Ability) ToFnSign() string {
	var params []string
	for _, param := range ability.FormalParams {
		params = append(params, param.Name)
	}
	return fmt.Sprintf("%s(%s)", ability.Name, strings.Join(params, ","))
}

func (ability *Ability) ToFnDesc() string {
	var desc strings.Builder
	desc.WriteString("【函数")
	desc.WriteString(ability.ToFnSign())
	desc.WriteString("提供")
	desc.WriteString(ability.Description)
	desc.WriteString(",")
	for _, param := range ability.FormalParams {
		desc.WriteString("参数")
		desc.WriteString(param.Name)
		desc.WriteString("是")
		desc.WriteString(param.Description)
		desc.WriteString(",")
	}
	result, _ := strings.CutSuffix(desc.String(), ",")
	return result + "】"
}
