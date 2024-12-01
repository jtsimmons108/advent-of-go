package main

import (
	"fmt"

	"simmons.com/advent-of-go/utils"
	"simmons.com/advent-of-go/y2024"
)

func main() {
	Run(y2024.Days[1])

}

func Run(d utils.Day) {
	fmt.Printf("Part1: %s\n", d.SolvePart1())
	fmt.Printf("Part2: %s\n", d.SolvePart2())
}