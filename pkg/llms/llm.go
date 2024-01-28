package llms

import (
	"context"
)

type Module interface {
	// GenerateContent generates content from a prompt.
	GenerateContent(ctx context.Context, prompt string) (string, error)
}
