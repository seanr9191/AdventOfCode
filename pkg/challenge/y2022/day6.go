package y2022

import (
	"go.uber.org/zap"

	"github.com/seanr9191/AdventOfCode/pkg/io/file"
)

type Day6 struct {
	Year      int
	Day       int
	InputFile string
	Logger    *zap.SugaredLogger
}

type System struct {
	signal     string
	markerSize int
}

func (s *System) FindFirstMarker() int {
	for i := range s.signal {
		foundData := make(map[rune]bool)
		for p := 0; p < s.markerSize; p++ {
			curr := rune(s.signal[i+p])
			if foundData[curr] {
				break
			}
			foundData[curr] = true
		}
		if len(foundData) == s.markerSize {
			return i + s.markerSize
		}
	}

	return -1
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

	system := System{signal: lines[0], markerSize: 4}
	return system.FindFirstMarker(), nil
}

func (d *Day6) Part2() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	system := System{signal: lines[0], markerSize: 14}
	return system.FindFirstMarker(), nil
}
