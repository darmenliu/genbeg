package parser

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"
)

type SourceFile struct {
	FileName string
	FileContent string
}

type SourceFileDict struct {
	SourceFiles map[string]SourceFile
}

func (s *SourceFileDict) AddSourceFile(fileName string, fileContent string) {
	s.SourceFiles[fileName] = SourceFile{FileName: fileName, FileContent: fileContent}
}

func (s *SourceFileDict) GetSourceFile(fileName string) (SourceFile, error) {
	file, ok := s.SourceFiles[fileName]
	if !ok {
		return SourceFile{}, fmt.Errorf("file not found")
	}
	return file, nil
}

func NewSourceFileDict() *SourceFileDict {
	return &SourceFileDict{SourceFiles: make(map[string]SourceFile)}
}

type CodeParser interface {
	ParseCode(text string) (SourceFileDict, error)
}

type GoCodeParser struct {
}

func NewGoCodeParser() *GoCodeParser {
	return &GoCodeParser{}
}

// ParseCode function Parse the code from markdown blocks and return a SourceFileDict
func (g *GoCodeParser) ParseCode(text string) (*SourceFileDict, error) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	sourceFileDict := NewSourceFileDict()

	// Regex to match code blocks

	regex := regexp.MustCompile("(\\s+)\n\\s*```[^\\n]*\n(.*?)```")

	// Find all matches
	matches := regex.FindAllStringSubmatch(text, -1)

	for _, match := range matches {
		// Get filename and content
		path := match[1] 
		content := match[2]

		// Clean filename
		path = regexp.MustCompile(`[\:<>"|?*]`).ReplaceAllString(path, "")
		path = regexp.MustCompile(`^\[(.*)\]$`).ReplaceAllString(path, "$1")
		path = regexp.MustCompile(`^`+(path)+`$`).ReplaceAllString(path, "$1")
		path = regexp.MustCompile(`\[]$`).ReplaceAllString(path, "")

		logger.Info("Adding file to source file dict", "path", path, "content", content)
		// Add to map
		sourceFileDict.AddSourceFile(path, content)
	}

	return sourceFileDict, nil
}
