package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/esheffield/adventofcode2018/utils"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide the path to the input as the first argument")
		os.Exit(1)
	}

	input := args[1]

	lines, err := utils.ReadFile(input)

	if err != nil {
		panic(err)
	}

	fmt.Println(lines[0])

	source := "[1518-08-30 00:03] Guard #1307 begins shift"
	line := strings.SplitAfterN(source, "]", 2)

	fmt.Println(line)

	re := regexp.MustCompile("#[\\d]+")

	fmt.Printf("%q\n", re.FindString(line[1]))

	t, err := time.Parse("[2006-01-02 15:04]", line[0])

	if err != nil {
		panic(err)
	}

	fmt.Println(t)
}
