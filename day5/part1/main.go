package main

import (
	"fmt"
	"os"

	"github.com/esheffield/adventofcode2018/utils"
	"github.com/golang-collections/collections/stack"
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

	diff := int32('a' - 'A')

	savedChars := stack.New()

	lastChar := int32(0)
	for _, char := range lines[0] {
		if savedChars.Len() == 0 {
			savedChars.Push(char)
		} else {
			lastChar = savedChars.Peek().(int32)
			curDiff := utils.Abs(lastChar - int32(char))
			if curDiff != diff {
				savedChars.Push(char)
			} else {
				savedChars.Pop()
			}
		}
	}

	fmt.Println(savedChars.Len())

	result := ""

	strLen := savedChars.Len()

	for i := 0; i < strLen; i++ {
		result = string(savedChars.Pop().(rune)) + result
	}

	fmt.Println("Result: ", result)
}
