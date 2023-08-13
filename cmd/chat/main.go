package main

import (
    "context"
    "fmt"
    "os"
    openai "github.com/sashabaranov/go-openai"
	"github.com/joho/godotenv"
)

func main() {
	if godotenv.Load() != nil {
		fmt.Println("Error loading .env file")
	}

    apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        fmt.Println("Please set your OPENAI_API_KEY.")
        return
    }
	client := openai.NewClient(apiKey)

    fmt.Print("Enter your prompt: ")
    var prompt string
    _, err := fmt.Scan(&prompt)
    if err != nil {
        fmt.Println("Failed to read prompt:", err)
        return
    }

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
