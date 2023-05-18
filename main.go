package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	// read .env file
	// read env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	for {
		// Create a new scanner to read from standard input
		scanner := bufio.NewScanner(os.Stdin)

		// Print a prompt to the terminal
		fmt.Print("Prompt (or 'file' to load from prompt.txt, or 'exit' to quit): ")

		// Scan the input until a newline character is encountered
		scanner.Scan()

		fmt.Println("Querying...")

		// Retrieve the input from the scanner
		input := scanner.Text()

		// Check if the user wants to exit
		if input == "exit" {
			break // Exit the loop
		}

		if input == "file" {
			// load the contents of prompt.txt
			file, err := os.Open("prompt.txt")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				input = scanner.Text()
			}
		}

		// Use input to call GPT 4
		client := openai.NewClient(os.Getenv("OPEN_AI_API_TOKEN"))
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT4,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: input,
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

	fmt.Println("Bye!")
}
