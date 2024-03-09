package main

import (
	"app/gpt"
	"app/gpt/entities"
	"log"
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

	log.Println(response.GetAnswer())
}
