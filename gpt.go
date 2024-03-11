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

	config := openai.DefaultConfig("token")
	config.BaseURL = baseURL
	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
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
		return "接口抽风了，再试下吧"
	}
	return resp.Choices[0].Message.Content
}
