package y2022

import (
	"strings"

	"go.uber.org/zap"

	"github.com/seanr9191/AdventOfCode/pkg/io/file"
)

type Day2 struct {
	Year      int
	Day       int
	InputFile string
	Logger    *zap.SugaredLogger
}

type battle struct {
	opposing move
	personal move
}

type move string

func (m move) GetWinningMove() move {
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

func (m move) GetLosingMove() move {
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

var matchingMoves = map[move]move{
	move("X"): move("A"),
	move("Y"): move("B"),
	move("Z"): move("C"),
}

func (b *battle) ChoiceScore() int {
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

func (b *battle) ResultScore() int {
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
		battle := &battle{
			opposing: move(moves[0]),
			personal: matchingMoves[move(moves[1])],
		}

		totalScore += battle.ChoiceScore() + battle.ResultScore()

		// d.Logger.Infof("Opponent: %v, You: %v, Choice Score: %v, Result Score: %v, Total Score: %v", moves[0], matchingMoves[move(moves[1])], battle.ChoiceScore(), battle.ResultScore(), totalScore)
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

		opponentMove := move(moves[0])
		var actualMove move
		switch moves[1] {
		case "X":
			actualMove = opponentMove.GetWinningMove()
		case "Y":
			actualMove = opponentMove
		case "Z":
			actualMove = opponentMove.GetLosingMove()
		}

		battle := &battle{
			opposing: opponentMove,
			personal: actualMove,
		}

		totalScore += battle.ChoiceScore() + battle.ResultScore()

		// d.Logger.Infof("battle %v, Opponent: %v, You: %v, Choice Score: %v, Result Score: %v, Total Score: %v", i, moves[0], moves[1], battle.ChoiceScore(), battle.ResultScore(), totalScore)
	}

	return totalScore, nil
}
