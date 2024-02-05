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

func FailureExit() {
	os.Exit(1)
}

func main() {

	ctx := context.Background()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("starting the application")
	model, err := gemini.NewGemini(ctx)
	if err != nil {
		logger.Error("failed to create model,", "err", err.Error())
		FailureExit()
	}
	defer model.CloseBackend()

	// Create nuwa-engineer workspace
	workspaceManager := NewDefaultWorkSpaceManager()
	err = workspaceManager.CreateWorkspace()
	if err != nil {
		logger.Error("failed to create workspace,", "err", err.Error())
		FailureExit()
	}

	// Create Golang project dir in the workspace, and initialize the project
	err = workspaceManager.CreateGolangProject("password_checker")
	if err != nil {
		logger.Error("failed to create golang project,", "err", err.Error())
		FailureExit()
	}

	// Initialize the Golang project
	err = workspaceManager.InitGolangProject("password_checker", "#PasswordChecker\n\n A password checker tool")
	if err != nil {
		logger.Error("failed to initialize golang project,", "err", err.Error())
		FailureExit()
	}

	logger.Info("model created, and sending request to generate content")

	userPrompt := `Please write a password checker, this tool help users to check if their
password is strong enough. The password should be at least 8 characters long, contain 
at least one uppercase letter, one lowercase letter,one number, and one special character. 
The tool should return a boolean value indicating whether the password is strong enough.`

	prompt := prompts.GetUserPrompt(userPrompt)

	fmt.Println(userPrompt)

	resp, err := model.GenerateContent(ctx, prompt)
	if err != nil {
		logger.Error("failed to generate content", "err", err.Error())
		FailureExit()
	}

	codeblocks, err := parser.NewGoCodeParser().ParseCode(resp)
	if err != nil {
		logger.Error("failed to parse code", "err", err.Error())
		FailureExit()
	}

	// print the response
	//fmt.Println(resp)

	for _, code := range codeblocks {
		code.ParseFileName()
		code.ParseFileContent()
		logger.Info("debug", "file name", code.FileName)
		fmt.Println(code.FileName)
		fmt.Println(code.FileContent)
	}

}
