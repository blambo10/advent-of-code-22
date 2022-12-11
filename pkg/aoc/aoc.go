package aoc

type AOC struct {
	ElvesCalories map[int]int
	SPRMatchData  []string
}

func New() *AOC {
	a := AOC{}

	return &a
}

func (a *AOC) Day(day string) error {

	switch day {
	case "1":
		inputData := make(map[int][]string)

		aggregateData, err := getDayData(day)

		if err != nil {
			return err
		}

		segregateCalories := processCalories(*aggregateData)

		inputData = segregateCalories

		caloriesByElf := GetElfCaloriesByElf(inputData)

		a.ElvesCalories = caloriesByElf
	case "2":
		aggregateData, err := getDayData(day)

		if err != nil {
			return err
		}

		SPRData := processSPRData(aggregateData)

		a.SPRMatchData = SPRData
	}

	return nil
}
