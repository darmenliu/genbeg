package workspace

import (
	"os"
)

// GetWorkspacePath retrieves the workspace path from the NUWA_WORKSPACE environment variable.
// If the environment variable is not set, an empty string is returned.
func GetWorkspacePath() string {
	return os.Getenv("NUWA_WORKSPACE")
}
