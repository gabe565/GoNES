package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gabe565/gones/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	output := "./docs"

	if err := os.RemoveAll(output); err != nil {
		log.Fatal(fmt.Errorf("failed to remove existing dir: %w", err))
	}

	if err := os.MkdirAll(output, 0o755); err != nil {
		log.Fatal(fmt.Errorf("failed to mkdir: %w", err))
	}

	rootCmd := cmd.New()
	if err := doc.GenMarkdownTree(rootCmd, output); err != nil {
		log.Fatal(fmt.Errorf("failed to generate markdown: %w", err))
	}
}
