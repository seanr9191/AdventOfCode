package y2022

import (
	"sort"
	"strconv"

	"go.uber.org/zap"

	"github.com/seanr9191/AdventOfCode/pkg/io/file"
)

type Day1 struct {
	Year      int
	Day       int
	InputFile string
	Logger    *zap.SugaredLogger
}

func (d *Day1) Solve() error {
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

func (d *Day1) Part1() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	var calories []int
	totalCalories := 0
	for _, line := range lines {
		if len(line) == 0 {
			calories = append(calories, totalCalories)
			totalCalories = 0
		} else {
			currCalories, err := strconv.Atoi(line)
			if err != nil {
				return nil, err
			}

			totalCalories += currCalories
		}
	}

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	totalCalories = calories[0]
	return totalCalories, nil
}

func (d *Day1) Part2() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	var calories []int
	totalCalories := 0
	for _, line := range lines {
		if len(line) == 0 {
			calories = append(calories, totalCalories)
			totalCalories = 0
		} else {
			currCalories, err := strconv.Atoi(line)
			if err != nil {
				return nil, err
			}

			totalCalories += currCalories
		}
	}

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	totalCalories = calories[0] + calories[1] + calories[2]
	return totalCalories, nil
}
