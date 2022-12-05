package y2022

import (
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

func peekFirst(stack []string) string {
	return stack[0]
}

func shift(stack []string) ([]string, string) {
	crate := stack[0]
	return stack[1:], crate
}

func shiftMany(stack []string, num int) ([]string, []string) {
	crates := stack[0:num]
	return stack[num:], crates
}

func prepend(stack []string, crate string) []string {
	stack = append(stack, "")
	copy(stack[1:], stack)
	stack[0] = crate
	return stack
}

func prependMany(stack []string, crates []string, reverse bool) []string {
	if reverse {
		crates = reverseSlice(crates)
	}

	newStack := make([]string, len(crates)+len(stack))
	copy(newStack[:len(crates)], crates)
	copy(newStack[len(crates):], stack)
	return newStack
}

func reverseSlice(stack []string) []string {
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}
	return stack
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
			line = strings.ReplaceAll(line, "move ", "")
			line = strings.ReplaceAll(line, " from ", ",")
			line = strings.ReplaceAll(line, " to ", ",")
			pieces := strings.Split(line, ",")
			numToMove, err := strconv.Atoi(pieces[0])
			if err != nil {
				return nil, err
			}
			from, err := strconv.Atoi(pieces[1])
			if err != nil {
				return nil, err
			}
			to, err := strconv.Atoi(pieces[2])
			if err != nil {
				return nil, err
			}

			s, crates := shiftMany(stacks[from-1], numToMove)
			stacks[from-1] = s
			p := prependMany(stacks[to-1], crates, true)
			stacks[to-1] = p
			moves++
			if moves%10000 == 0 {
				d.Logger.Infof("Completed move %v.", moves)
			}
		}
	}

	result := ""
	for i := 0; i < len(stacks); i++ {
		result += peekFirst(stacks[i])
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
			// Moves start on 10
			line = strings.ReplaceAll(line, "move ", "")
			line = strings.ReplaceAll(line, " from ", ",")
			line = strings.ReplaceAll(line, " to ", ",")
			pieces := strings.Split(line, ",")
			numToMove, err := strconv.Atoi(pieces[0])
			if err != nil {
				return nil, err
			}
			from, err := strconv.Atoi(pieces[1])
			if err != nil {
				return nil, err
			}
			to, err := strconv.Atoi(pieces[2])
			if err != nil {
				return nil, err
			}

			s, crates := shiftMany(stacks[from-1], numToMove)
			stacks[from-1] = s
			stacks[to-1] = prependMany(stacks[to-1], crates, false)
			moves++
			if moves%10000 == 0 {
				d.Logger.Infof("Completed move %v.", moves)
			}
		}
	}

	result := ""
	for i := 0; i < len(stacks); i++ {
		result += peekFirst(stacks[i])
	}

	return result, nil
}
