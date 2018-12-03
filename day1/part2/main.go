package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func index(arr []int, elt int) int {
	for i, v := range arr {
		if v == elt {
			return i
		}
	}

	return -1
}

func main() {
	var changes []int
	var seen []int
	freq := 0
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide the path to the input as the first argument")
		os.Exit(1)
	}

	input := args[1]

	fmt.Println("File: ", input)

	file, err := os.Open(input)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		offset, _ := strconv.Atoi(str)
		changes = append(changes, offset)
	}

	fmt.Printf("There are %d changes.\n", len(changes))

	fmt.Println("Initial index: ", index(seen, freq))

	for i := 0; index(seen, freq) == -1; i++ {
		seen = append(seen, freq)
		if i == len(changes) {
			i = 0
		}
		freq += changes[i]
	}

	fmt.Printf("Frequency seen twice: %d\n", freq)
}
