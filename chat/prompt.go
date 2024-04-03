package chat

type Prompt string

const (
	Translate = "当我输入中文时你需要翻译成英文，我可能输入单个汉字、词语、成语、词组、短语、短句、长句、俗语、文档等等，你的职责仅仅是翻译为合适的英文并回答给我。" +
		"当我输入英文时你需要翻译成中文,当我输入中文时你需要翻译成英文，我可能输入单词、词组、短语、短句、长句、俗语、文档等等，你的职责仅仅是翻译并回答给我。" +
		"不需要解释原本的输入，不需要输出分析过程，直接输出翻译结果。当我需要你解释的时候我会以[详细翻译:]作为前缀，详细翻译时 你应该以词典的形式给出翻译，包括：翻译、拼音或音标、例句、其他解释等。"
)

func StrToPrompt(promptStr string) Prompt {
	return Prompt(promptStr)
}

func GetFunctionCall(abilitiesJson string) Prompt {
	return Prompt("下面是由函数对象(func)组成的JSON数组(funcArray):\n```json\n" +
		abilitiesJson +
		"\n```\nfunc每个字段的含义是:\n```json\n{\"id\":\"函数标识\",\"name\":\"函数名\",\"desc\":\"函数说明\",\"params\":[{\"name\":\"形参名称\",\"desc\":\"形参说明\"}]}\n```\n" +
		"你需要从输入的自然语言中推测意图(intent),并根据函数的desc到funcArray进行匹配,请仔细分析函数的desc它是你准确匹配的关键.如果能批量成功则根据形参的desc提取参数,请仔细分析形参的desc它是你准确提取实参的关键.\n" +
		" - 匹配成功则返回一个经过minify处理的JSON文本:\n ```json\n{\"status\":\"1\",\"funcId\":\"函数id\",\"params\":[\"从自然语言中提取的实参Array\"]}\n ```\n" +
		" - 匹配失败则返回一个经过minify处理的JSON文本:\n ```json\n{\"status\":\"0\",\"guide\":\"如果funcArray中存在与intent接近的func则需要输出为能引导后续输入以达到完全匹配func目的引导语此时你需要继续从后面的用户输入获取信息直到能完全匹配上func或者直到输入了其他能完全匹配func的语言，否则输出为空\"}\n ```\n" +
		"下面是一个对话示例来让你更好的理解上面的要求:\n" +
		" - intent匹配func\n" +
		"   input:明天早上8点叫我起床\n" +
		"   output:{\"status\":\"1\",\"funcId\":\"ability.SetPeriodicReminds\",,\"params\":[\"115200\",\"该起床了\"]}\n" +
		" - intent接近的func\n" +
		"   input:一会提醒我喝水\n" +
		"   output:{\"status\":\"0\",guide:\"5分钟后提醒你可以吗?\"}\n" +
		"   input:可以\n" +
		"   output:{\"status\":\"1\",\"funcId\":\"ability.SetDisposableReminds\",\"params\":[\"1\",\"300\",\"喝水时间带到了，去喝口水吧\"]}\n" +
		" - intent完全不能匹配func\n" +
		"   input:我饿了\n" +
		"   output:{\"status\":\"0\",guide:\"\"}")
}
