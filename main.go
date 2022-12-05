package main

import (
	"fmt"
	"github.com/blambo10/advent-of-code-22/pkg/aoc"
)

func main() {
	aoc := aoc.New()
	_ = aoc

	for i, inputs := range aoc.DayInputs {
		day := i + 1
		fmt.Println("Day: %s", day)
		fmt.Println("Data: %s", *inputs)
		fmt.Println("\n")
	}

	fmt.Println("adfa")
}
