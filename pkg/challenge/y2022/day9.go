package y2022

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"go.uber.org/zap"

	"github.com/seanr9191/AdventOfCode/pkg/io/file"
)

type Day9 struct {
	Year      int
	Day       int
	InputFile string
	Logger    *zap.SugaredLogger
}

type simulation struct {
	tailPositions map[string]int
	rope          rope
}

func newSimulation(knotCount int) *simulation {
	knots := make([]*coordinate, 0)
	for i := 0; i < knotCount; i++ {
		knots = append(knots, &coordinate{
			X: 0,
			Y: 0,
		})
	}
	sim := &simulation{
		tailPositions: make(map[string]int),
		rope: rope{
			Knots: knots,
		},
	}
	sim.tailPositions[sim.rope.Knots[knotCount-1].String()] = 1
	return sim
}

func (s *simulation) Simulate(direction string, steps int) {
	for i := 0; i < steps; i++ {
		// Move the head
		switch direction {
		case "U":
			s.rope.Knots[0].Y++
		case "D":
			s.rope.Knots[0].Y--
		case "L":
			s.rope.Knots[0].X--
		case "R":
			s.rope.Knots[0].X++
		}

		for n := 1; n < len(s.rope.Knots); n++ {
			currKnot := s.rope.Knots[n-1]
			nextKnot := s.rope.Knots[n]

			// If touching, continue
			if math.Abs(float64(currKnot.Y-nextKnot.Y)) <= 1 &&
				math.Abs(float64(currKnot.X-nextKnot.X)) <= 1 {
				continue
			}

			if currKnot.X > nextKnot.X {
				nextKnot.X++
			} else if currKnot.X < nextKnot.X {
				nextKnot.X--
			}

			if currKnot.Y > nextKnot.Y {
				nextKnot.Y++
			} else if currKnot.Y < nextKnot.Y {
				nextKnot.Y--
			}

			if n == len(s.rope.Knots)-1 {
				s.tailPositions[nextKnot.String()] = 1
			}
		}
	}
}

func (s *simulation) TailSpots() int {
	return len(s.tailPositions)
}

type rope struct {
	Knots []*coordinate
}

type coordinate struct {
	X int
	Y int
}

func (c *coordinate) String() string {
	return fmt.Sprintf("%v,%v", c.X, c.Y)
}

func (d *Day9) Solve() error {
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

func (d *Day9) Part1() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	simulation := newSimulation(2)
	for _, line := range lines {
		pieces := strings.Split(line, " ")
		direction := pieces[0]
		steps, err := strconv.Atoi(pieces[1])
		if err != nil {
			return nil, err
		}
		simulation.Simulate(direction, steps)
	}

	return simulation.TailSpots(), nil
}

func (d *Day9) Part2() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	simulation := newSimulation(10)
	for _, line := range lines {
		pieces := strings.Split(line, " ")
		direction := pieces[0]
		steps, err := strconv.Atoi(pieces[1])
		if err != nil {
			return nil, err
		}
		simulation.Simulate(direction, steps)
	}

	return simulation.TailSpots(), nil
}
