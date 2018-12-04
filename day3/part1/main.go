package main

import (
	"fmt"
	"log"
	"os"

	"github.com/esheffield/adventofcode2018/utils"
)

type Patch struct {
	id      int
	xOffset int
	yOffset int
	width   int
	height  int
}

func processPatch(patch string) Patch {
	var id, xOffset, yOffset, width, height int
	fmt.Sscanf(patch, "#%d @ %d,%d: %dx%d", &id, &xOffset, &yOffset, &width, &height)
	p := Patch{id, xOffset, yOffset, width, height}
	return p
}

func main() {
	path := os.Args[1]
	lines, err := utils.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println(lines)

	fmt.Println(processPatch(lines[0]))
}
