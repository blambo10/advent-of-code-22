package aoc

type AOC struct {
	DayInputs map[int]map[int][]string
}

func New() *AOC {
	a := AOC{}
	a.GetInputData()

	return &a
}

func (a *AOC) GetInputData() error {

	inputData := make(map[int]map[int][]string)

	days := 25
	day := 1
	for day <= days {
		aggregateData, err := getInputs(1)

		if err != nil {
			continue
		}

		segregateData := processInputs(*aggregateData)

		inputData[day] = segregateData

		day++
	}

	a.DayInputs = inputData

	return nil
}
