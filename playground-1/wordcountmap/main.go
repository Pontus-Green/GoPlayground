package main

import (
	"cmp"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: wordcount <file>")
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

	words := strings.Fields(text)
	wordCount := map[string]int{}

	for _, v := range words {
		wordCount[v] += 1
	}

	var sorted []Pair
	for k, v := range wordCount {
		sorted = append(sorted, Pair{k, v})
	}

	slices.SortFunc(sorted, func(a, b Pair) int {
		return cmp.Compare(b.Value, a.Value)
	})

	for _, pair := range sorted {
		fmt.Printf("%s: %d\n", pair.Key, pair.Value)
	}

}
