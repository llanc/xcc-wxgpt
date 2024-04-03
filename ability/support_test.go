package ability

import (
	"fmt"
	"testing"
)

func TestRegister(t *testing.T) {
	abilityList := Register("support.go")
	str := abilityList.ToString()
	fmt.Println(str)
}
