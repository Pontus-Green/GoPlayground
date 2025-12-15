package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func printHelp() {
	fmt.Println("Usage: wordcount file --optional_args")
	fmt.Println("If optional args is not provided wordcount will return")
	fmt.Println("Lines count: <Number of lines in the file>")
	fmt.Println("Words count: <Number of words in the file>")
	fmt.Println("characters count: <number of characters in the file>")
	fmt.Println("optional_args:")
	fmt.Println("--lines: will only return the lines count")
	fmt.Println("--words: will only return the words count")
	fmt.Println("--characters: will only return the characters count")
}

func main() {
	validArgs := []string{"--help", "--lines", "--words", "--characters"}

	if len(os.Args) < 2 {
		fmt.Println("Usage: wordcount <file>")
		fmt.Println("For additional info: wordcount --help")
		os.Exit(1)
	}
	if slices.Contains(os.Args, "--help") {
		printHelp()
		os.Exit(0)
	}

	for _, v := range os.Args[2:] {
		if !slices.Contains(validArgs, v) {
			fmt.Println("Invalid argument:", v)
			printHelp()
			os.Exit(1)
		}
		fmt.Println("validargs", validArgs)
	}

	fileName := os.Args[1]

	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	fmt.Println("Proccesing wordcount for file...", fileName)

	text := string(data)
	if len(os.Args) == 2 {
		lines := strings.Count(text, "\n") + 1
		words := len(strings.Fields(text))
		characters := len(data)

		fmt.Println("Lines count:", lines)
		fmt.Println("Words count:", words)
		fmt.Println("characters count:", characters)
		os.Exit(0)
	}

	if slices.Contains(os.Args, "--lines") {
		lines := strings.Count(text, "\n") + 1
		fmt.Println("Lines count:", lines)
	}
	if slices.Contains(os.Args, "--words") {
		words := len(strings.Fields(text))
		fmt.Println("Words count:", words)
	}
	if slices.Contains(os.Args, "--characters") {
		characters := len(data)
		fmt.Println("characters count:", characters)
	}
	os.Exit(0)
}
