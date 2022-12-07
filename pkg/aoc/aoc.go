package aoc

type AOC struct {
	ElvesCalories map[int]int
}

func New() *AOC {
	a := AOC{}
	a.Init()

	return &a
}

func (a *AOC) Init() error {

	inputData := make(map[int][]string)

	aggregateData, err := getCalories()

	if err != nil {
		return err
	}

	segregateCalories := processCalories(*aggregateData)

	inputData = segregateCalories

	caloriesByElf := GetElfCaloriesByElf(inputData)

	a.ElvesCalories = caloriesByElf

	return nil
}
