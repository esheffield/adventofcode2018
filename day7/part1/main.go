package main

import (
	"fmt"
	"os"

	"github.com/esheffield/adventofcode2018/utils"
)

const offset = rune('A')

type Step struct {
	id   rune
	next []int
	prev []int
}

func parseLines(lines []string) []Step {
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

	return steps
}

func removeElt(elts []int, elt int) []int {
	loc := -1
	for i, e := range elts {
		if e == elt {
			loc = i
			break
		}
	}

	if loc != -1 {
		return append(elts[:loc], elts[loc+1:]...)
	}

	return elts
}

func findRune(runes []rune, r rune) int {
	for i, elt := range runes {
		if elt == r {
			return i
		}
	}

	return -1
}

func doSteps(steps []Step) []rune {
	var seq []rune
	stepCnt := 0
	for _, step := range steps {
		if step.id != rune(-1) {
			stepCnt++
		}
	}

	for len(seq) != stepCnt {
		fmt.Println("Seq: ", string(seq))
		for i, step := range steps {
			if step.id != rune(-1) && findRune(seq, step.id) == -1 && len(step.prev) == 0 {
				for _, nextStepIdx := range step.next {
					nextStep := steps[nextStepIdx]
					nextStep.prev = removeElt(nextStep.prev, i)
					steps[nextStepIdx] = nextStep
				}
				seq = append(seq, step.id)
				break
			}
		}
	}

	return seq
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

	steps := parseLines(lines)

	fmt.Println(steps)

	path := doSteps(steps)

	fmt.Println(string(path))
}
