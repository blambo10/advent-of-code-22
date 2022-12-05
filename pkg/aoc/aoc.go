package aoc

type AOC struct {
	DayInputs map[int]*string
}

func New() *AOC {
	a := AOC{}
	a.GetInputData()

	return &a
}

func (a *AOC) GetInputData() error {

	inputData := make(map[int]*string)

	days := 25
	day := 1
	for day <= days {
		data, err := GetInputs(1)

		if err != nil {
			continue
		}

		inputData[day] = data

		day++
	}

	a.DayInputs = inputData

	return nil
}
