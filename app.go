package main

import (
	"app/gpt"
	"app/gpt/entities"
	"os"
	"strings"
	"github.com/joho/godotenv"
)

func main() {
	// Loading environment params
	godotenv.Load(".env")

	response := gpt.Ask([]entities.Message{
		entities.Message{
			Role:    "user",
			Content: strings.Join(os.Args[1:], " "),
		},
	})

	println("\n - " + response.GetAnswer() + "\n")
}
