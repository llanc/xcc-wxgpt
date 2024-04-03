package chat

import (
	"os"
	"testing"
	"xcc-wxgpt/ability"
)

func TestPromptChat(t *testing.T) {
	os.Setenv("BASE_URL", "")
	os.Setenv("TOKEN", "")
	os.Setenv("MODEL", "gpt-4")
	gpt := &Chat{}
	gpt.Init()

	abilities := ability.Register("../ability/support.go")
	responseMsg := gpt.PromptChat(GetFunctionCall(abilities.ToString()), "明天下午2点我要去趟超市", 0.1)
	funcCall, err := GetFuncCallFromGpt(responseMsg)
	if err != nil {
		println(responseMsg)
	}
	funcCall.Local()

}
