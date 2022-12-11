package aoc

import (
	"strings"
)

const (
	A, X = 1, 1
	B, Y = 2, 2
	C, Z = 3, 3

	draw   = 3
	loser  = 0
	winner = 6
)

func processSPRData(s *string) []string {
	return strings.Split(*s, "\n")
}

//func initialgetMatchScore(matchData string) (int, int) {
//	var p1score int
//	var p2score int
//
//	mScore := strings.Split(matchData, " ")
//
//	p1play := mScore[0]
//	p2play := mScore[1]
//	switch p1play {
//	case "A":
//		switch p2play {
//		case "Z":
//			p1score = A + winner
//			p2score = Z
//		case "Y":
//			p1score = A
//			p2score = Y + winner
//		case "X":
//			p1score = A + draw
//			p2score = X + draw
//		}
//	case "B":
//		switch p2play {
//		case "Z":
//			p1score = B
//			p2score = Z + winner
//		case "Y":
//			p1score = B + draw
//			p2score = Y + draw
//		case "X":
//			p1score = B + winner
//			p2score = X
//		}
//	case "C":
//		switch p2play {
//		case "Z":
//			p1score = C + draw
//			p2score = Z + draw
//		case "Y":
//			p1score = C + winner
//			p2score = Y
//		case "X":
//			p1score = C
//			p2score = X + winner
//		}
//	}
//
//	return p1score, p2score
//}

func getMatchScore(matchData string) int {
	var score int

	mScore := strings.Split(matchData, " ")

	//X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"
	p1play := mScore[0]
	outcome := mScore[1]

	switch p1play {
	case "A":
		switch outcome {
		case "Z":
			score = B + winner
		case "Y":
			score = A + draw
		case "X":
			score = C
		}
	case "B":
		switch outcome {
		case "Z":
			score = C + winner
		case "Y":
			score = B + draw
		case "X":
			score = A
		}
	case "C":
		switch outcome {
		case "Z":
			score = A + winner
		case "Y":
			score = C + draw
		case "X":
			score = B
		}
	}

	return score
}

func GetTotalScore(sprData []string) int {
	var total int

	for _, matchData := range sprData {
		if len(matchData) > 0 {
			score := getMatchScore(matchData)

			total += score

		}
	}

	return total
}
