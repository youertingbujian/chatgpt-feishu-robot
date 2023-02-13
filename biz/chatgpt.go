package biz

import (
	"chatgpt/config"
	"context"
	"github.com/sashabaranov/go-gpt3"
	"github.com/sirupsen/logrus"
)

// ChatGPT ...
func ChatGPT(msg string) string {
	client := gogpt.NewClient(config.Conf.OpenAiApiKey)
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:       gogpt.GPT3TextDavinci003,
		Prompt:      msg,
		Temperature: 0,
		MaxTokens:   4000,
	}
	resp, err := client.CreateCompletion(ctx, req)
	if err != nil {
		logrus.WithError(err).Errorf("ChatGPT接口异常")
		return "ChatGPT接口异常"
	}
	return resp.Choices[0].Text
}
