package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
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
		freq += offset
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("Frequency: %d\n", freq)
}
