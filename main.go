package main

import (
	"fmt"
	"github.com/blambo10/advent-of-code-22/pkg/aoc"
	"os"
	"sort"
)

func main() {
	args := os.Args[0:]

	a := aoc.New()

	switch args[1] {
	case "1":
		var cal []int
		highestCalories := 0

		for _, calories := range a.ElvesCalories {
			if calories > highestCalories {
				highestCalories = calories
			}
		}

		fmt.Println(highestCalories)

		for _, calories := range a.ElvesCalories {
			cal = append(cal, calories)
		}

		sort.Ints(cal)

		first := cal[len(cal)-1]
		second := cal[len(cal)-2]
		third := cal[len(cal)-3]

		fmt.Println(first + second + third)
	case "2":

	}
}
