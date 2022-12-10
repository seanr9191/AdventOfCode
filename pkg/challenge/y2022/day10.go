package y2022

import (
	"fmt"
	"strconv"
	"strings"

	"go.uber.org/zap"

	"github.com/seanr9191/AdventOfCode/pkg/io/file"
)

type Day10 struct {
	Year      int
	Day       int
	InputFile string
	Logger    *zap.SugaredLogger
}

func (d *Day10) Solve() error {
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

type cpu struct {
	Register   int
	CycleCount int
}

type state struct {
	Cycle    int
	Register int
}

func (c *cpu) cycle(instruction string) (int, error) {
	c.CycleCount++

	var value int
	var err error
	if strings.HasPrefix(instruction, "addx") {
		c.CycleCount++
		parts := strings.Split(instruction, " ")
		value, err = strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}
		c.Register += value
	}

	return value, err
}

func (d *Day10) Part1() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	totalStrength := 0
	cpu := cpu{Register: 1, CycleCount: 0}
	for _, instruction := range lines {
		registerMovement, err := cpu.cycle(instruction)
		if err != nil {
			return nil, err
		}

		if cpu.CycleCount%40 == 0 || (cpu.CycleCount-1)%40 == 0 {
			continue
		}

		cyclesToCheck := []*state{{
			Cycle:    cpu.CycleCount,
			Register: cpu.Register - registerMovement,
		}}
		if strings.HasPrefix(instruction, "addx") {
			cyclesToCheck = append(cyclesToCheck, &state{
				Cycle:    cpu.CycleCount - 1,
				Register: cpu.Register - registerMovement,
			})
		}

		for _, state := range cyclesToCheck {
			if state.Cycle%20 == 0 {
				strength := state.Cycle * state.Register
				totalStrength += strength
				break
			}
		}
	}

	return totalStrength, nil
}

func (d *Day10) Part2() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	totalStrength := 0
	cpu := cpu{Register: 2, CycleCount: 0}
	for _, instruction := range lines {
		registerMovement, err := cpu.cycle(instruction)
		if err != nil {
			return nil, err
		}

		var cyclesToCheck []*state
		if strings.HasPrefix(instruction, "addx") {
			cyclesToCheck = append(cyclesToCheck, &state{
				Cycle:    cpu.CycleCount - 1,
				Register: cpu.Register - registerMovement,
			})
		}
		cyclesToCheck = append(cyclesToCheck, &state{
			Cycle:    cpu.CycleCount,
			Register: cpu.Register - registerMovement,
		})

		for _, state := range cyclesToCheck {

			if state.Cycle%40 == 0 {
				fmt.Print("\n")
				continue
			}

			if state.Cycle%40 >= state.Register-1 && state.Cycle%40 <= state.Register+1 {
				fmt.Print("# ")
			} else {
				fmt.Print(". ")
			}
		}
	}

	return totalStrength, nil
}
