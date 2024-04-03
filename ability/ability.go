package ability

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"strings"
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

func (list *Abilities) ToString() string {
	jsonData, err := json.Marshal(list)
	if err != nil {
		fmt.Printf("Error marshaling to JSON: %v\n", err)
		return ""
	}
	return string(jsonData)
}

func Register(filePath string) *Abilities {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	pkg := &ast.Package{
		Name:  "ability",
		Files: make(map[string]*ast.File),
	}
	pkg.Files["support.go"] = f

	docPkg := doc.New(pkg, "", 0)

	var list Abilities
	for i, fun := range docPkg.Funcs {
		if fun.Name == "register" {
			continue
		}
		docs := strings.Split(fun.Doc, "\n")
		if i == len(docs) {
			continue
		}
		funcDesc, _ := strings.CutPrefix(docs[0], fun.Name+" ")
		ability := Ability{
			ID:          pkg.Name + "." + fun.Name,
			Name:        fun.Name,
			Description: funcDesc,
		}
		for i, paramsDoc := range docs {
			if i == 0 || i+1 == len(docs) {
				continue
			}
			paramDoc := strings.Split(paramsDoc, " ")
			paramName := paramDoc[0]
			paramDesc := strings.Join(paramDoc[1:], " ")
			param := Param{
				Name:        paramName,
				Description: paramDesc,
			}
			ability.FormalParams = append(ability.FormalParams, param)
		}
		list = append(list, ability)
	}
	return &list
}

var AbilityMap = map[string]func(args ...string){
	"ability.SetPeriodicReminds":   SetPeriodicReminds,
	"ability.SetDisposableReminds": SetDisposableReminds,
}
