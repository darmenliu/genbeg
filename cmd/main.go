package main

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {

	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text("Write a story about a magic backpack."))
	if err != nil {
		log.Fatal(err)
	}

	// print the first candidate
	log.Print(resp.Candidates[0].Content.Parts[0])
	for _, candidate := range resp.Candidates {
		log.Print(candidate.Content.Parts[0])
		for _, part := range candidate.Content.Parts {
			log.Print(part)
		}
	}

}
