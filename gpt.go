package main

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
	"strings"
)

func chat(msg string) string {
	token := os.Getenv("TOKEN")
	if len(strings.TrimSpace(token)) == 0 {
		panic("TOKEN IS BLANK")
	}
	baseURL := os.Getenv("BASE_URL")
	if len(strings.TrimSpace(baseURL)) == 0 {
		panic("baseURL IS BLANK")
	}
	model := os.Getenv("MODEL")
	if len(strings.TrimSpace(baseURL)) == 0 {
		println("MODEL IS BLANK SET DEFAULT GPT3.5")
		model = "gpt-3.5-turbo"
	}

	config := openai.DefaultConfig(token)
	config.BaseURL = baseURL
	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: msg,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "上游接口出问题了，稍后再试下吧^_^"
	}
	return resp.Choices[0].Message.Content
}
