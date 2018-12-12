package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/esheffield/adventofcode2018/utils"
)

type Point struct {
	x int32
	y int32
}

type Area struct {
	point Point
	area  int32
}

type Plot struct {
	owner    int
	root     bool
	distance int32
}

type Region struct {
	xOffset int32
	yOffset int32
	width   int32
	height  int32
	plots   [][]Plot
}

func parsePoints(lines []string) []Area {
	var areas []Area
	for _, line := range lines {
		coords := strings.Split(line, ", ")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		areas = append(areas, Area{Point{int32(x), int32(y)}, 0})
	}

	return areas
}

func manhattenDistance(from Point, to Point) int32 {
	return utils.Abs(from.x-to.x) + utils.Abs(from.y-to.y)
}

func findMinMax(areas []Area) (Point, Point) {
	if len(areas) == 0 {
		return Point{}, Point{}
	}
	first := areas[0].point
	minX := first.x
	minY := first.y
	maxX := first.x
	maxY := first.y

	for _, area := range areas {
		point := area.point
		if point.x < minX {
			minX = point.x
		}
		if point.y < minY {
			minY = point.y
		}
		if point.x > maxX {
			maxX = point.x
		}
		if point.y > maxY {
			maxY = point.y
		}
	}

	return Point{minX, minY}, Point{maxX, maxY}
}

func getRegionDimensions(minXY Point, maxXY Point) (int32, int32) {
	regionWidth := maxXY.x - minXY.x + 3
	regionHeight := maxXY.y - minXY.y + 3

	return regionWidth, regionHeight
}

func initPlots(regionWidth int32, regionHeight int32) [][]Plot {
	region := make([][]Plot, regionHeight)
	for i := range region {
		region[i] = make([]Plot, regionWidth)
	}

	return region
}

func createRegion(areas []Area) Region {
	minXY, maxXY := findMinMax(areas)
	width, height := getRegionDimensions(minXY, maxXY)
	return Region{
		xOffset: minXY.x - 1,
		yOffset: minXY.y - 1,
		width:   width,
		height:  height,
		plots:   initPlots(width, height),
	}
}

func initCoords(region Region, areas []Area) {
	plots := region.plots
	for i, area := range areas {
		point := area.point
		plots[point.y-region.yOffset][point.x-region.xOffset] = Plot{i, true, 0}
	}
	region.plots = plots
}

func computeDistances(region Region, areas []Area) {
	plots := region.plots

	for y, row := range plots {
		for x, plot := range row {
			if plot.root {
				continue
			}
			plot.distance = int32(region.height + region.width)
			curPoint := Point{int32(x), int32(y)}
			for i, area := range areas {
				point := area.point
				distance := manhattenDistance(point, curPoint)
				if distance < plot.distance {
					plot.distance = distance
					plot.owner = i
				} else if distance == plot.distance {
					plot.distance = distance
					plot.owner = -1
				}
			}
			row[x] = plot
		}
	}

	region.plots = plots
}

func computeAreas(region Region, areas []Area) {
	plots := region.plots
	maxHeight := int(region.height - 1)
	maxWidth := int(region.width - 1)

	for y, row := range plots {
		for x, plot := range row {
			if plot.owner == -1 || areas[plot.owner].area == -1 {
				continue
			} else if x == 0 || x == maxWidth || y == 0 || y == maxHeight {
				// on the edge, so parent has infinite area
				areas[plot.owner] = Area{areas[plot.owner].point, -1}
			} else {
				areas[plot.owner] = Area{areas[plot.owner].point, areas[plot.owner].area + 1}
			}
		}
	}
}

func printRegion(region Region) {
	for _, row := range region.plots {
		for _, plot := range row {
			if plot.root {
				fmt.Printf("*%1s* ", string(plot.owner+'A'))
			} else if plot.owner == -1 {
				fmt.Printf(" .  ")
			} else {
				fmt.Printf(" %1s  ", string(plot.owner+'a'))
			}
		}
		fmt.Println()
	}
}

func printAreas(areas []Area) {
	for i, area := range areas {
		fmt.Printf("%3d\t%d\n", i, area.area)
	}
}

func findLargestArea(areas []Area) Area {
	var maxArea Area
	for _, area := range areas {
		if area.area > maxArea.area {
			maxArea = area
		}
	}

	return maxArea
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

	areas := parsePoints(lines)

	region := createRegion(areas)

	initCoords(region, areas)

	computeDistances(region, areas)

	computeAreas(region, areas)

	// printRegion(region)

	// printAreas(areas)

	largestArea := findLargestArea(areas)

	fmt.Println("Largest area: ", largestArea)
}
