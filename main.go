package main

import (
	"fmt"
	"github.com/blambo10/advent-of-code-22/pkg/aoc"
	"os"
	"sort"
)

func main() {
	args := os.Args[0:]
	day := args[1]

	a := aoc.New()

	switch day {
	case "1":
		a.Day(day)
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
		err := a.Day(day)

		if err != nil {
			fmt.Println(err)
		}

		totalScore := aoc.GetTotalScore(a.SPRMatchData)

		fmt.Println(totalScore)
	}
}
