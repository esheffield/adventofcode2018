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

type LogEntry struct {
	t        time.Time
	guardNum int
	action   Action
}

type Guard struct {
	guardNum   int
	minutes    [60]int
	totalSlept int
}

func parseLine(source string) LogEntry {
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

	return LogEntry{t, guardNum, action}
}

func findMaxSleepGuard(guards map[int]Guard) (Guard, int) {
	var selectedGuard Guard
	maxMinute := -1
	timesAsleep := -1

	for _, guard := range guards {
		mostSleptMinute := findMostSleptMinute(guard)

		if asleep := guard.minutes[mostSleptMinute]; asleep > timesAsleep {
			selectedGuard = guard
			maxMinute = mostSleptMinute
			timesAsleep = asleep
		}
	}

	return selectedGuard, maxMinute
}

func findMostSleptMinute(guard Guard) int {
	max := -1
	maxMinute := -1

	for i, timesSlept := range guard.minutes {
		if timesSlept > max {
			max = timesSlept
			maxMinute = i
		}
	}

	return maxMinute
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

	fmt.Println("Parsing entries...")
	var entries []LogEntry
	for _, line := range lines {
		entries = append(entries, parseLine(line))
	}

	fmt.Println("Sorting...")
	sort.Slice(entries[:], func(i, j int) bool {
		return entries[i].t.Before(entries[j].t)
	})

	guards := make(map[int]Guard)

	curGuard := -1
	var sleepTime time.Time

	fmt.Println("Processing entries...")
	for _, entry := range entries {
		switch entry.action {
		case BeginShift:
			curGuard = entry.guardNum
			if _, ok := guards[curGuard]; !ok {
				guards[curGuard] = Guard{guardNum: curGuard}
			}
		case Sleep:
			sleepTime = entry.t
		case Wake:
			for minute := sleepTime.Minute(); minute < entry.t.Minute(); minute++ {
				guard := guards[curGuard]
				guard.minutes[minute]++
				guard.totalSlept++
				guards[curGuard] = guard
			}
		}
	}

	fmt.Println("Find guard with most slept minute...")
	sleepiestGuard, mostSleptMinute := findMaxSleepGuard(guards)

	fmt.Println("Guard with most slept minute: ", sleepiestGuard)
	fmt.Println("Most slept minute: ", mostSleptMinute)

	fmt.Println("Code: ", sleepiestGuard.guardNum*mostSleptMinute)
}
