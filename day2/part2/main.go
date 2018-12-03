package main

import (
	"fmt"
	"log"
	"os"

	"github.com/esheffield/adventofcode2018/utils"
)

func findNearDups(ids []string) (string, int) {
	for i := range ids[0] {
		counts := make(map[string]int)
		for _, id := range ids {
			str := id[0:i] + id[i+1:]
			counts[str]++
			if counts[str] == 2 {
				return str, i
			}
		}
	}

	return "", -1
}

func main() {
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

	common, index := findNearDups(lines)

	fmt.Printf("Found: %s differs at %d\n", common, index)
}
