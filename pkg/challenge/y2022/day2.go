package y2022

import (
	"AdventOfCode/pkg/io/file"
	"go.uber.org/zap"
	"strings"
)

type Day2 struct {
	Year      int
	Day       int
	InputFile string
	Logger    *zap.SugaredLogger
}

type Battle struct {
	opposing Move
	personal Move
}

type Move string

func (m Move) GetWinningMove() Move {
	switch m {
	case "A":
		return "C"
	case "B":
		return "A"
	case "C":
		return "B"
	}

	return ""
}

func (m Move) GetLosingMove() Move {
	switch m {
	case "A":
		return "B"
	case "B":
		return "C"
	case "C":
		return "A"
	}

	return ""
}

var matchingMoves = map[Move]Move{
	Move("X"): Move("A"),
	Move("Y"): Move("B"),
	Move("Z"): Move("C"),
}

func (b *Battle) ChoiceScore() int {
	switch b.personal {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	}

	return 0
}

func (b *Battle) ResultScore() int {
	if b.personal.GetWinningMove() == b.opposing {
		return 6
	} else if b.personal.GetLosingMove() == b.opposing {
		return 0
	}

	return 3
}

func (d *Day2) Solve() error {
	a1, err := d.Part1()
	if err != nil {
		return err
	}
	a2, err := d.Part2()
	if err != nil {
		return err
	}

	d.Logger.Infof("Day %v, %v completed. Part 1: %v, Part 2: %v", d.Day, d.Year, a1, a2)

	return nil
}

func (d *Day2) Part1() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	totalScore := 0
	for _, line := range lines {
		moves := strings.Split(line, " ")
		battle := &Battle{
			opposing: Move(moves[0]),
			personal: matchingMoves[Move(moves[1])],
		}

		totalScore += battle.ChoiceScore() + battle.ResultScore()

		// d.Logger.Infof("Opponent: %v, You: %v, Choice Score: %v, Result Score: %v, Total Score: %v", moves[0], matchingMoves[Move(moves[1])], battle.ChoiceScore(), battle.ResultScore(), totalScore)
	}

	return totalScore, nil
}

func (d *Day2) Part2() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	totalScore := 0
	for _, line := range lines {
		moves := strings.Split(line, " ")

		opponentMove := Move(moves[0])
		var actualMove Move
		switch moves[1] {
		case "X":
			actualMove = opponentMove.GetWinningMove()
		case "Y":
			actualMove = opponentMove
		case "Z":
			actualMove = opponentMove.GetLosingMove()
		}

		battle := &Battle{
			opposing: opponentMove,
			personal: actualMove,
		}

		totalScore += battle.ChoiceScore() + battle.ResultScore()

		// d.Logger.Infof("Battle %v, Opponent: %v, You: %v, Choice Score: %v, Result Score: %v, Total Score: %v", i, moves[0], moves[1], battle.ChoiceScore(), battle.ResultScore(), totalScore)
	}

	return totalScore, nil
}
