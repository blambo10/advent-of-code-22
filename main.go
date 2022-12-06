package main

import (
	"fmt"
	"github.com/blambo10/advent-of-code-22/pkg/aoc"
)

func main() {
	a := aoc.New()

	for day, data := range a.DayInputs {

		switch day {
		case 1:
			elfCalories := aoc.ProcessDayOne(data)

			for elf, calories := range elfCalories {
				fmt.Println("Elf: ", elf)
				fmt.Println("Calories: ", calories)
				fmt.Println("\n")
			}

		}
		//data
		//day := i + 1
		//fmt.Println("Day: %s", day)
		//fmt.Println("Data: %s", *inputs)
		//fmt.Println("\n")
	}

	fmt.Println("adfa")
}
