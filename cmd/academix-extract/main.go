// Command academix-extract extracts Academix embedded assets to a directory.
//
// Usage:
//
//	academix-extract [output-dir]
//
// Default output: ~/.reasonix/academix-assets
package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"reasonix/internal/skill/builtincontent"
)

func main() {
	outDir := filepath.Join(os.Getenv("HOME"), ".reasonix", "academix-assets")
	if len(os.Args) > 1 {
		outDir = os.Args[1]
	}

	assets := builtincontent.AcademixAssets()
	count := 0

	err := fs.WalkDir(assets, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return os.MkdirAll(filepath.Join(outDir, p), 0755)
		}
		data, err := fs.ReadFile(assets, p)
		if err != nil {
			return err
		}
		dest := filepath.Join(outDir, p)
		if err := os.MkdirAll(filepath.Dir(dest), 0755); err != nil {
			return err
		}
		if err := os.WriteFile(dest, data, 0644); err != nil {
			return err
		}
		count++
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Extracted %d files to %s\n", count, outDir)
}
