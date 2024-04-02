package ability

import (
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"strconv"
	"strings"
)

// SetPeriodicReminds 提供设置固定周期任务的能力
// time 固定周期,是一个大于0的整数秒数
// message 具体提醒内容
func SetPeriodicReminds(args ...string) {
	//time := args[0]
	//message := args[1]
}

// SetDisposableReminds 提供设置一次性提醒的能力
// timeType 时间类型:取值1时代表相对0点的时间,取值2时代表相对当前时间的时间
// time 与type想对应，代表相对0点或者相对当前时间偏移的任意整数秒数
// message 具体提醒内容
func SetDisposableReminds(args ...string) {
	//timeType := args[0]
	//time := args[1]
	//message := args[2]
}

// Register 将当前文件中所有非Register函数转为Ability并加入到List返回
func Register() Abilities {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "support.go", nil, parser.ParseComments)
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
			ID:          strconv.Itoa(i + 1),
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
	return list
}

var AbilityMap = map[string]func(args ...string){
	"1": SetPeriodicReminds,
	"2": SetDisposableReminds,
}
