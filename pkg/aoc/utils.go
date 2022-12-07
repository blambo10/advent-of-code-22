package aoc

import (
	"fmt"
	"github.com/blambo10/advent-of-code-22/pkg/config"
	"github.com/blambo10/advent-of-code-22/pkg/utils"
	"net/http"
	"strconv"
	"strings"
)

func getCalories() (*string, error) {

	//var data interface{}
	cfg := config.GetAOCConfig()
	client := utils.GetRestyClient().R()
	url := fmt.Sprintf("https://%s/2022/day/%s/input", cfg.URL, "1")

	resp, err := client.SetCookie(&http.Cookie{
		Name:  "session",
		Value: cfg.Cookie,
	}).Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() >= 400 {
		return nil, fmt.Errorf("an error occured fetching inputs")
	}

	data := string(resp.Body())

	return &data, nil
}

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
