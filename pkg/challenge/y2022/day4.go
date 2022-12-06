package y2022

import (
	"strconv"
	"strings"

	"go.uber.org/zap"

	"github.com/seanr9191/AdventOfCode/pkg/data_structure/boundary"
	"github.com/seanr9191/AdventOfCode/pkg/io/file"
)

func parseBounds(input string) (int, int, error) {
	parts := strings.Split(input, "-")
	lower, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	upper, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}

	return lower, upper, nil
}

type Day4 struct {
	Year      int
	Day       int
	InputFile string
	Logger    *zap.SugaredLogger
}

func (d *Day4) Solve() error {
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

func (d *Day4) Part1() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	encompassed := 0
	for _, line := range lines {
		bounds := strings.Split(line, ",")
		l1, u1, err := parseBounds(bounds[0])
		if err != nil {
			return nil, err
		}
		l2, u2, err := parseBounds(bounds[1])
		if err != nil {
			return nil, err
		}

		firstElf := &boundary.Boundary{
			Lower: l1,
			Upper: u1,
		}

		secondElf := &boundary.Boundary{
			Lower: l2,
			Upper: u2,
		}

		if firstElf.EncompassedBy(secondElf) || firstElf.Encompasses(secondElf) {
			encompassed++
		}
	}

	return encompassed, nil
}

func (d *Day4) Part2() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	overlaps := 0
	for _, line := range lines {
		bounds := strings.Split(line, ",")
		l1, u1, err := parseBounds(bounds[0])
		if err != nil {
			return nil, err
		}
		l2, u2, err := parseBounds(bounds[1])
		if err != nil {
			return nil, err
		}

		firstElf := &boundary.Boundary{
			Lower: l1,
			Upper: u1,
		}

		secondElf := &boundary.Boundary{
			Lower: l2,
			Upper: u2,
		}

		if firstElf.Overlaps(secondElf) {
			overlaps++
		}
	}

	return overlaps, nil
}
