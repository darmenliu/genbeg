package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"nuwa-engineer/pkg/llms/gemini"
	"nuwa-engineer/pkg/parser"
	"nuwa-engineer/pkg/prompts"
)

func main() {

	ctx := context.Background()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("starting the application")
	model, err := gemini.NewGemini(ctx)
	if err != nil {
		logger.Error("failed to create model,", err.Error())
		return
	}
	defer model.CloseBackend()

	logger.Info("model created, and sending request to generate content")

	userPrompt := `Please write a password checker, this tool help users to check if their
password is strong enough. The password should be at least 8 characters long, contain 
at least one uppercase letter, one lowercase letter,one number, and one special character. 
The tool should return a boolean value indicating whether the password is strong enough.`

	prompt := prompts.GetUserPrompt(userPrompt)

	fmt.Println(userPrompt)

	resp, err := model.GenerateContent(ctx, prompt)
	if err != nil {
		logger.Error("failed to generate content", err.Error())
		return
	}

	codeblocks, err := parser.NewGoCodeParser().ParseCode(resp)
	if err != nil {
		logger.Error("failed to parse code", err.Error())
		return
	}

	// print the response
	//fmt.Println(resp)

	for _, code := range codeblocks {
		code.ParseFileName()
		code.ParseFileContent()
		logger.Info("File Name:", code.FileName)
		fmt.Println(code.FileName)
		fmt.Println(code.FileContent)
	}

}
