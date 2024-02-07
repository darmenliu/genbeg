package main

import (
	"context"
	"fmt"
	"os"

	"nuwa-engineer/pkg/llms/gemini"

	goterm "github.com/c-bata/go-prompt"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)


func GenerateContent(ctx context.Context, prompt string) (string, error) {
	ctx = context.Background()
	model, err := gemini.NewGemini(ctx)
	if err != nil {
		return "", fmt.Errorf("Failed to create model: %w", err)
	}
	defer model.CloseBackend()

	resp, err := model.GenerateContent(ctx, prompt)
	if err != nil {
		return "", fmt.Errorf("Failed to generate content: %w", err)
	}
	return resp, nil
}

func FailureExit() {
	os.Exit(1)
}

func executor(in string) {
	fmt.Println("You: " + in)
	if in == "" {
		return
	}
	ctx := context.Background()
	rsp, err := GenerateContent(ctx, in)
	if err != nil {
		fmt.Println("NUWA: " + err.Error())
		return
	}
	fmt.Println("NUWA: " + rsp)
}

func completer(in goterm.Document) []goterm.Suggest {
	s := []goterm.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
		{Text: "groups", Description: "Combine users with specific rules"},
	}
	return goterm.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}


func main() {

	// Initialize a big text display with the letters "Nuwa" and "Engineer"
	// "P" is displayed in cyan and "Term" is displayed in light magenta
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("Nuwa", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle(" Engineer", pterm.FgLightMagenta.ToStyle())).
		Render() // Render the big text to the terminal

	//logger := slog.New(slog.NewTextHandler(os.Stdout, nil))


	// // Create nuwa-engineer workspace
	// workspaceManager := NewDefaultWorkSpaceManager()
	// err = workspaceManager.CreateWorkspace()
	// if err != nil {
	// 	logger.Error("failed to create workspace,", "err", err.Error())
	// 	FailureExit()
	// }

// 	// Create Golang project dir in the workspace, and initialize the project
// 	err = workspaceManager.CreateGolangProject("password_checker")
// 	if err != nil {
// 		logger.Error("failed to create golang project,", "err", err.Error())
// 		FailureExit()
// 	}

// 	// Initialize the Golang project
// 	err = workspaceManager.InitGolangProject("password_checker", "#PasswordChecker\n\n A password checker tool")
// 	if err != nil {
// 		logger.Error("failed to initialize golang project,", "err", err.Error())
// 		FailureExit()
// 	}

// 	logger.Info("model created, and sending request to generate content")

// 	userPrompt := `Please write a password checker, this tool help users to check if their
// password is strong enough. The password should be at least 8 characters long, contain 
// at least one uppercase letter, one lowercase letter,one number, and one special character. 
// The tool should return a boolean value indicating whether the password is strong enough.`

// 	prompt := prompts.GetUserPrompt(userPrompt)

// 	fmt.Println(userPrompt)

// 	resp, err := model.GenerateContent(ctx, prompt)
// 	if err != nil {
// 		logger.Error("failed to generate content", "err", err.Error())
// 		FailureExit()
// 	}

// 	codeblocks, err := parser.NewGoCodeParser().ParseCode(resp)
// 	if err != nil {
// 		logger.Error("failed to parse code", "err", err.Error())
// 		FailureExit()
// 	}

// 	// print the response
// 	//fmt.Println(resp)

// 	for _, code := range codeblocks {
// 		code.ParseFileName()
// 		code.ParseFileContent()
// 		logger.Info("debug", "file name", code.FileName)
// 		fmt.Println(code.FileName)
// 		fmt.Println(code.FileContent)
// 	}

	p := goterm.New(
		executor,
		completer,
		goterm.OptionPrefix(">>> "),
		goterm.OptionTitle("NUWA ENGINEER"),
	)
	p.Run()

}
