package main

import (
	"fmt"
	"go-reloaded/textprocessor"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
		os.Exit(1)
	}

	result := textprocessor.Process(string(data))

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output file: %v\n", err)
		os.Exit(1)
	}
}
