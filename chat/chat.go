package chat

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
	"strings"
)

type Chat struct {
	model  string
	client *openai.Client
}

func (chat *Chat) Init() {
	baseURL := os.Getenv("BASE_URL")
	if len(strings.TrimSpace(baseURL)) == 0 {
		panic("baseURL IS BLANK")
	}
	token := os.Getenv("TOKEN")
	if len(strings.TrimSpace(token)) == 0 {
		panic("TOKEN IS BLANK")
	}
	model := os.Getenv("MODEL")
	if len(strings.TrimSpace(baseURL)) == 0 {
		println("MODEL IS BLANK SET DEFAULT GPT3.5")
		model = "gpt-3.5-turbo"
	}
	chat.model = model
	config := openai.DefaultConfig(token)
	config.BaseURL = baseURL
	client := openai.NewClientWithConfig(config)
	chat.client = client
}

func (chat *Chat) RawChat(msg string) string {
	resp, err := chat.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: chat.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: msg,
				},
			},
		},
	)

	if err != nil {
		return chat.errorHandle(err)
	}
	return resp.Choices[0].Message.Content
}

func (chat *Chat) SimpleChat(prompt Prompt, msg string) string {
	return chat.PromptChat(prompt, msg, 0.7)
}

func (chat *Chat) StrictChat(prompt Prompt, msg string) string {
	return chat.PromptChat(prompt, msg, 0.1)
}

func (chat *Chat) PromptChat(prompt Prompt, msg string, temperature float32) string {
	resp, err := chat.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       chat.model,
			Temperature: temperature,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: string(prompt),
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: msg,
				},
			},
		},
	)

	if err != nil {
		return chat.errorHandle(err)
	}
	return resp.Choices[0].Message.Content
}

func (chat *Chat) errorHandle(err error) string {
	fmt.Printf("ChatCompletion error: %v\n", err)
	return "上游接口出问题了，稍后再试下吧^_^"
}
