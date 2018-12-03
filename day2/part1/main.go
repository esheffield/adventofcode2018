package main

import (
	"fmt"
	"log"
	"os"

	"github.com/esheffield/adventofcode2018/utils"
)

func checkID(id string) (bool, bool) {
	hasDouble, hasTriple := false, false
	counts := make(map[rune]int)

	for _, ch := range id {
		counts[ch] = counts[ch] + 1
	}

	for _, count := range counts {
		if count == 2 {
			hasDouble = true
		}
		if count == 3 {
			hasTriple = true
		}
	}

	return hasDouble, hasTriple
}

func main() {
	doubles, triples := 0, 0
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide the path to the input as the first argument")
		os.Exit(1)
	}

	input := args[1]

	fmt.Println("File: ", input)

	lines, err := utils.ReadLines(input)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for _, str := range lines {
		hasDouble, hasTriple := checkID(str)
		if hasDouble {
			doubles++
		}
		if hasTriple {
			triples++
		}
	}

	fmt.Printf("Doubles: %d\tTriples: %d\tChecksum: %d\n", doubles, triples, doubles*triples)
}
