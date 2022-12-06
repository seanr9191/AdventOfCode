package y2022

import (
	"AdventOfCode/pkg/data_structure/slice"
	"AdventOfCode/pkg/io/file"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

type Day5 struct {
	Year      int
	Day       int
	InputFile string
	Logger    *zap.SugaredLogger
}

func (d *Day5) Solve() error {
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

func (d *Day5) Part1() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	stacks := make([][]string, 9)
	moves := 0
	for _, line := range lines {
		cleanLine := strings.ReplaceAll(line, " ", "")
		if strings.HasPrefix(cleanLine, "[") {
			// Reading crates
			// line has len of 3 per crate/stack. Need to parse 3 characters at a time
			runeLine := []rune(line)
			for loc, ind := 0, 0; loc < len(line); loc, ind = loc+4, ind+1 {
				chunk := string(runeLine[loc : loc+3])
				if strings.HasPrefix(chunk, "[") {
					stacks[ind] = append(stacks[ind], string(chunk[1]))
				}
			}
		} else if strings.HasPrefix(cleanLine, "move") {
			// Processing moves
			pieces := strings.Split(line, " ")
			numToMove, err := strconv.Atoi(pieces[1])
			if err != nil {
				return nil, err
			}
			from, err := strconv.Atoi(pieces[3])
			if err != nil {
				return nil, err
			}
			to, err := strconv.Atoi(pieces[5])
			if err != nil {
				return nil, err
			}

			s, crates := slice.ShiftMany(stacks[from-1], numToMove)
			stacks[from-1] = s
			crates = slice.Reverse(crates)
			p := slice.PrependMany(stacks[to-1], crates)
			stacks[to-1] = p
			moves++
			if moves%10000 == 0 {
				d.Logger.Infof("Completed move %v.", moves)
			}
		}
	}

	result := ""
	for i := 0; i < len(stacks); i++ {
		result += slice.PeekFirst(stacks[i])
	}

	return result, nil
}

func (d *Day5) Part2() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	stacks := make([][]string, 9)
	moves := 0
	for _, line := range lines {
		cleanLine := strings.ReplaceAll(line, " ", "")
		if strings.HasPrefix(cleanLine, "[") {
			// Reading crates
			// line has len of 3 per crate/stack. Need to parse 3 characters at a time
			runeLine := []rune(line)
			for loc, ind := 0, 0; loc < len(line); loc, ind = loc+4, ind+1 {
				chunk := string(runeLine[loc : loc+3])
				if strings.HasPrefix(chunk, "[") {
					stacks[ind] = append(stacks[ind], string(chunk[1]))
				}
			}
		} else if strings.HasPrefix(cleanLine, "move") {
			// Processing moves
			pieces := strings.Split(line, " ")
			numToMove, err := strconv.Atoi(pieces[1])
			if err != nil {
				return nil, err
			}
			from, err := strconv.Atoi(pieces[3])
			if err != nil {
				return nil, err
			}
			to, err := strconv.Atoi(pieces[5])
			if err != nil {
				return nil, err
			}

			s, crates := slice.ShiftMany(stacks[from-1], numToMove)
			stacks[from-1] = s
			stacks[to-1] = slice.PrependMany(stacks[to-1], crates)
			moves++
			if moves%10000 == 0 {
				d.Logger.Infof("Completed move %v.", moves)
			}
		}
	}

	result := ""
	for i := 0; i < len(stacks); i++ {
		result += slice.PeekFirst(stacks[i])
	}

	return result, nil
}
