package main

import (
	"fmt"
	"os"

	"github.com/esheffield/adventofcode2018/utils"
	"github.com/golang-collections/collections/stack"
)

const diff = int32('a' - 'A')

func reactPolymer(polymer string, ignore rune) string {
	savedChars := stack.New()

	lastChar := int32(0)
	for _, char := range polymer {
		if char == ignore || char == ignore+diff {
			continue
		}
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

	result := ""
	strLen := savedChars.Len()

	for i := 0; i < strLen; i++ {
		result = string(savedChars.Pop().(rune)) + result
	}

	return result
}

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

	minChar := rune(0)
	minLen := len(lines[0]) + 1

	for ch := 'A'; ch <= 'Z'; ch++ {
		result := reactPolymer(lines[0], ch)
		if len(result) < minLen {
			minLen = len(result)
			minChar = ch
		}
	}

	fmt.Println("Min length: ", minLen, " by omitting ", string(minChar))
}
