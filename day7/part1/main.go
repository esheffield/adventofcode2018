package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/esheffield/adventofcode2018/utils"
)

const offset = rune('A')

type Step struct {
	id   rune
	next []int
	prev []int
}

func parseLines(lines []string) ([]Step, int) {
	steps := make([]Step, 26)
	for i := 0; i < 26; i++ {
		step := steps[i]
		step.id = rune(-1)
		steps[i] = step
	}

	for _, line := range lines {
		chars := []rune(line)
		stepId := chars[5]
		nextStepId := chars[36]
		step := steps[stepId-offset]
		nextStep := steps[nextStepId-offset]
		if nextStep.id != nextStepId {
			nextStep.id = nextStepId
		}
		nextStep.prev = append(nextStep.prev, int(stepId-offset))
		steps[nextStepId-offset] = nextStep
		step.id = stepId
		step.next = append(step.next, int(nextStepId-offset))
		steps[stepId-offset] = step
	}

	start := -1
	for i, step := range steps {
		if step.id != rune(-1) {
			sort.Ints(step.next)
			sort.Ints(step.prev)
			if len(step.prev) == 0 {
				start = i
			}
		}
	}

	return steps, start
}

func doSteps(steps []Step) []rune {
	var seq []rune
	for i, step := range steps {
		if step.id != rune(-1) && len(step.prev) == 0 {
			seq = append(seq, traverse(steps, i)...)
		}
	}

	return seq
}

func traverse(steps []Step, start int) []rune {
	var seq []rune
	step := steps[start]

	fmt.Println("Step: ", step)
	for _, nextStepId := range step.next {
		nextStep := steps[nextStepId]
		fmt.Println("\t", step.id, " -> Next: ", nextStep)
		if len(nextStep.prev) > 0 && nextStep.prev[len(nextStep.prev)-1] == start {
			seq = append(seq, traverse(steps, nextStepId)...)
		}
	}

	return append([]rune{step.id}, seq...)
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

	steps, start := parseLines(lines)

	fmt.Println("Start: ", start)
	fmt.Println(steps)

	path := doSteps(steps)

	fmt.Println(string(path))
}
