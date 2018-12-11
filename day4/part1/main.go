package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/esheffield/adventofcode2018/utils"
)

type Action int

const (
	None       Action = iota
	BeginShift Action = iota
	Sleep      Action = iota
	Wake       Action = iota
)

type Entry struct {
	t        time.Time
	guardNum int
	action   Action
}

func parseLine(source string) Entry {
	line := strings.SplitAfterN(source, "]", 2)

	re := regexp.MustCompile("#([\\d]+)")
	description := strings.TrimSpace(line[1])
	action := None
	guardNum := -1
	if strings.HasPrefix(description, "wakes") {
		action = Wake
	} else if strings.HasPrefix(description, "falls") {
		action = Sleep
	} else {
		action = BeginShift
		guardNumParts := re.FindStringSubmatch(description)
		if len(guardNumParts) == 2 {
			guardNum, _ = strconv.Atoi(guardNumParts[1])
		}
	}
	t, err := time.Parse("[2006-01-02 15:04]", line[0])

	if err != nil {
		panic(err)
	}

	return Entry{t, guardNum, action}
}

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

	var entries []Entry
	for _, line := range lines {
		entries = append(entries, parseLine(line))
	}

	sort.Slice(entries[:], func(i, j int) bool {
		return entries[i].t.Before(entries[j].t)
	})

	var sleepMinutes [][60]int

}
