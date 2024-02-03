package main

import (
	"fmt"
	"log/slog"
	"nuwa-engineer/pkg/dir"
	"nuwa-engineer/pkg/workspace"
	"os"
)

type WorkSpaceManager interface {
	// CreateWorkspace creates a new workspace.
	CreateWorkspace() error
}

type DefaultWorkSpaceManager struct{}

func NewDefaultWorkSpaceManager() WorkSpaceManager {
	return &DefaultWorkSpaceManager{}
}

func (d DefaultWorkSpaceManager) CreateWorkspace() error {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	workspacePath := workspace.GetWorkspacePath()
	if workspacePath == "" {
		logger.Error("failed to get workspace path")
		return fmt.Errorf("failed to get workspace path")
	}

	dirCreator := dir.NewDefaultDirectoryCreator()
	err := dirCreator.CreateDir(workspacePath)
	if err != nil {
		logger.Error("failed to create workspace", err.Error())
		return fmt.Errorf("failed to create workspace: %w", err)
	}

	return nil
}
