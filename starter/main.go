package main

import (
	"fmt"
	"os"

	"github.com/esheffield/adventofcode2018/utils"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide the path to the input as the first argument")
		os.Exit(1)
	}

	input := args[1]

	fmt.Println("Reading ", input)
	lines, err := utils.ReadFile(input)

	if err != nil {
		panic(err)
	}

	fmt.Println(lines)
}
