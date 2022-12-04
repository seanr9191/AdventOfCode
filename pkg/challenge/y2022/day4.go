package y2022

import (
	"AdventOfCode/pkg/io/file"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

type bound struct {
	Upper int
	Lower int
}

func newBound(input string) (*bound, error) {
	parts := strings.Split(input, "-")
	lower, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	upper, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	return &bound{
		Upper: upper,
		Lower: lower,
	}, nil
}

func (b *bound) Encompasses(other *bound) bool {
	return b.Lower <= other.Lower && b.Upper >= other.Upper
}

func (b *bound) Encompassed(other *bound) bool {
	return other.Lower <= b.Lower && other.Upper >= b.Upper
}

func (b *bound) Overlaps(other *bound) bool {
	if b.Encompassed(other) || b.Encompasses(other) {
		return true
	}

	if b.Lower <= other.Lower {
		return other.Lower <= b.Upper
	} else {
		return other.Upper >= b.Lower
	}
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
		first, err := newBound(bounds[0])
		if err != nil {
			return nil, err
		}
		second, err := newBound(bounds[1])
		if err != nil {
			return nil, err
		}

		if first.Encompassed(second) || first.Encompasses(second) {
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
		first, err := newBound(bounds[0])
		if err != nil {
			return nil, err
		}
		second, err := newBound(bounds[1])
		if err != nil {
			return nil, err
		}

		if first.Overlaps(second) {
			overlaps++
		}
	}

	return overlaps, nil
}
