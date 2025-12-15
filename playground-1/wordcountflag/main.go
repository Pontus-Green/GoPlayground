package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	linesFlag := flag.Bool("lines", false, "print only lines count")
	wordsFlag := flag.Bool("words", false, "print only words count")
	charactersFlag := flag.Bool("characters", false, "print only characters count")

	flag.Usage = func() {
		fmt.Println("Usage: wordcount <file> [--lines] [--words] [--characters]")
		fmt.Println("If no optional flags are provided, all counts are printed.")
	}

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	fileName := args[0]

	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	fmt.Println("Processing wordcount for file:", fileName)

	text := string(data)
	lines := strings.Count(text, "\n") + 1
	words := len(strings.Fields(text))
	characters := len(data)

	if !*linesFlag && !*wordsFlag && !*charactersFlag {
		*linesFlag, *wordsFlag, *charactersFlag = true, true, true
	}

	if *linesFlag {
		fmt.Println("Lines count:", lines)
	}
	if *wordsFlag {
		fmt.Println("Words count:", words)
	}
	if *charactersFlag {
		fmt.Println("Characters count:", characters)
	}
}
