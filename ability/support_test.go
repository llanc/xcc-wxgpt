package ability

import (
	"fmt"
	"testing"
)

func TestRegister(t *testing.T) {
	abilityList := Register()
	str, _ := abilityList.ToString()
	fmt.Println(str)
}
