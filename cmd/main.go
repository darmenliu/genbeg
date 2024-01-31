package main

import (
	"context"
	"log"

	"nuwa-engineer/pkg/llms/gemini"
	"nuwa-engineer/pkg/prompts"
)

func main() {

	ctx := context.Background()

	model, err := gemini.NewGemini(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer model.CloseBackend()

	resp, err := model.GenerateContent(ctx, "Write a story about a magic backpack.")
	if err != nil {
		log.Fatal(err)
	}

	// print the response
	log.Print(resp)
	log.Print(prompts.GetSysPrompt())
}
