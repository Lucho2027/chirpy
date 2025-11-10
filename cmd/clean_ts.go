package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run clean_ts_interfaces.go <path-to-directory>")
		os.Exit(1)
	}

	root := os.Args[1]
	err := filepath.Walk(root, processFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func processFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() || !strings.HasSuffix(info.Name(), ".ts") {
		return nil
	}

	lines, err := readLines(path)
	if err != nil {
		return fmt.Errorf("reading %s: %w", path, err)
	}

	var kept []string
	keepBlock := false

	for _, line := range lines {
		trim := strings.TrimSpace(line)

		// Start of an interface or type
		if strings.HasPrefix(trim, "export interface ") || strings.HasPrefix(trim, "export type ") {
			keepBlock = true
			kept = append(kept, line)
			continue
		}

		if keepBlock {
			kept = append(kept, line)
			if strings.HasPrefix(trim, "}") {
				keepBlock = false
			}
		}
	}

	if len(kept) == 0 {
		// If no interfaces/types, clear file
		err = os.WriteFile(path, []byte{}, 0644)
		if err != nil {
			return fmt.Errorf("clearing %s: %w", path, err)
		}
		fmt.Printf("🧹 Cleared (no types): %s\n", path)
		return nil
	}

	err = os.WriteFile(path, []byte(strings.Join(kept, "\n")+"\n"), 0644)
	if err != nil {
		return fmt.Errorf("writing %s: %w", path, err)
	}

	fmt.Printf("✅ Cleaned: %s\n", path)
	return nil
}

func readLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
