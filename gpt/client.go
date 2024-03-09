package gpt

import (
	"app/gpt/entities"
	"app/gpt/enum/endpoints"
	"app/gpt/enum/models"
	statuscodes "app/gpt/enum/status_codes"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

const openai_api_key = ""

func getOpenAiSecret() string {
	if openai_api_key == "" {
		return os.Getenv("OPENAI_API_KEY")
	}

	return openai_api_key
}

func Ask(messages []entities.Message) entities.ResponseInterface {

	requestBody := entities.Request{
		Model:    models.Gpt35Turbo,
		Messages: messages,
	}

	jsonBody, err := json.Marshal(requestBody)

	if err != nil {
		log.Fatal("Error marshalling json", err.Error())
	}

	request, _ := http.NewRequest(http.MethodPost, endpoints.Completions, bytes.NewBuffer(jsonBody))

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+getOpenAiSecret())

	response, err := (&http.Client{}).Do(request)

	if err != nil {
		log.Fatal("Asking error", err.Error())
	}

	defer response.Body.Close()

	if response.StatusCode == statuscodes.Unauthorized {
		log.Println("Unauthorized")
		os.Exit(response.StatusCode)
	}

    body, err := io.ReadAll(response.Body)

    if err != nil {
        panic(err.Error())
    }

    var responseModel entities.Response

    json.Unmarshal(body, &responseModel)

    return responseModel
}
