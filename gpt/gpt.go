// -----------------------------------------------------------------------
// Copyright (c) 2022 Akuvox Corporation and Akubela-Eevee Contributors.
// All rights reserved.
// -----------------------------------------------------------------------

package gpt

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

func Gpt(msg string) (string, error) {
	config := openai.DefaultConfig(os.Getenv("OPENAI_KEY"))
	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
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
		return "", err
	}
	fmt.Println(resp)
	return resp.Choices[0].Message.Content, nil
}
