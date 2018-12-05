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
	overlap bool
}

func processPatch(patch string) Patch {
	var id, xOffset, yOffset, width, height int
	fmt.Sscanf(patch, "#%d @ %d,%d: %dx%d", &id, &xOffset, &yOffset, &width, &height)
	p := Patch{id, xOffset, yOffset, width, height, false}
	return p
}

func main() {
	var patches []Patch
	var cloth [1000][1000][]int

	path := os.Args[1]
	lines, err := utils.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	multipatch := 0
	for _, line := range lines {
		patches = append(patches, processPatch(line))
	}

	for p, patch := range patches {
		for y := 0; y < patch.height; y++ {
			for x := 0; x < patch.width; x++ {
				row := patch.yOffset + y
				col := patch.xOffset + x
				cloth[row][col] = append(cloth[row][col], p)
			}
		}
	}

	for _, row := range cloth {
		for _, col := range row {
			if len(col) > 1 {
				for _, p := range col {
					patches[p].overlap = true
				}
				multipatch++
			}
		}
	}
	fmt.Println("Num shared patches: ", multipatch)
	for _, patch := range patches {
		if !patch.overlap {
			fmt.Println("Non-overlaping patch: ", patch)
		}
	}
}
