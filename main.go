package main

import (
	"fmt"
	"go-reloaded/transforms"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input> <output>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	text := string(data)
	lines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	var processed []string
	for _, line := range lines {
		processed = append(processed, transforms.Apply(line))
	}
	text = strings.Join(processed, "\n")

	err = os.WriteFile(outputFile, []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}
}
