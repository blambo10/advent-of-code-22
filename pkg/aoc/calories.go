package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

// func processCalories(s string) map[int][]string {
func processCalories(s string) map[int][]string {
	elves := make(map[int][]string)

	inputDataAggregate := strings.Split(s, "\n")

	elf := 0
	for _, inputData := range inputDataAggregate {

		if inputData == "" {
			elf++
			continue
		}

		elves[elf] = append(elves[elf], inputData)
	}

	return elves
}

func GetElfCaloriesByElf(elfData map[int][]string) map[int]int {

	elfCalories := make(map[int]int)

	for i, elf := range elfData {
		calculated := 0
		for _, data := range elf {
			calories, err := strconv.Atoi(data)

			if err != nil {
				fmt.Println(data)
				fmt.Println(err)
				continue
			}
			calculated += calories
		}

		elfCalories[i] = calculated
	}

	return elfCalories
}
