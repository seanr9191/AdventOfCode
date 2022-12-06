package y2022

import (
	"AdventOfCode/pkg/io/file"
	"go.uber.org/zap"
)

type Day6 struct {
	Year      int
	Day       int
	InputFile string
	Logger    *zap.SugaredLogger
}

func (d *Day6) Solve() error {
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

func (d *Day6) Part1() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		d.Logger.Info(line)
	}

	return 0, nil
}

func (d *Day6) Part2() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		d.Logger.Info(line)
	}

	return 0, nil
}
