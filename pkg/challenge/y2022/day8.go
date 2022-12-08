package y2022

import (
	"go.uber.org/zap"
	"strconv"

	"github.com/seanr9191/AdventOfCode/pkg/io/file"
)

type Day8 struct {
	Year      int
	Day       int
	InputFile string
	Logger    *zap.SugaredLogger
}

type forest struct {
	size    int
	trees   [][]int
	visible [][]int
	scenic  [][]int
}

func (f *forest) PlaceTrees(row []int, rowNum int) {
	copy(f.trees[rowNum], row)
}

func (f *forest) calculateVisible() int {
	if f.visible != nil && f.visible[0][0] == 1 {
		// We've calculated this before. Let's just count the visible.
		count := 0
		for i := range f.visible {
			for j := range f.visible[i] {
				if f.visible[i][j] == 1 {
					count++
				}
			}
		}
		return count
	} else {
		// We need to populate visible and then get the count.
		f.visible = make([][]int, f.size)
		for i := range f.visible {
			f.visible[i] = make([]int, f.size)
		}

		for i := range f.visible {
			for j := range f.visible[i] {
				if i == 0 || i == f.size-1 || j == 0 || j == f.size-1 {
					// This is the outer edge. These are always visible.
					f.visible[i][j] = 1
				}
			}
		}

		for i := range f.visible {
			for j := range f.visible[i] {
				if f.visible[i][j] == 0 {
					f.visible[i][j] = f.IsVisible(i, j)
				}
			}
		}

		return f.calculateVisible()
	}
}

func (f *forest) IsVisible(i, j int) int {
	good := true
	for y := i - 1; y >= 0; y-- {
		if f.trees[y][j] >= f.trees[i][j] {
			good = false
			break
		}
	}
	if good {
		return 1
	}

	good = true
	for y := i + 1; y <= f.size-1; y++ {
		if f.trees[y][j] >= f.trees[i][j] {
			good = false
			break
		}
	}
	if good {
		return 1
	}

	good = true
	for x := j - 1; x >= 0; x-- {
		if f.trees[i][x] >= f.trees[i][j] {
			good = false
			break
		}
	}
	if good {
		return 1
	}

	good = true
	for x := j + 1; x <= f.size-1; x++ {
		if f.trees[i][x] >= f.trees[i][j] {
			good = false
			break
		}
	}
	if good {
		return 1
	}

	return -1
}

func (f *forest) calculateScenic() int {
	if f.scenic != nil && f.scenic[1][1] >= 1 {
		// We've calculated this before. Let's just find the most scenic.
		mostScenic := 0
		for i := range f.scenic {
			for j := range f.scenic[i] {
				if f.scenic[i][j] > mostScenic {
					mostScenic = f.scenic[i][j]
				}
			}
		}
		return mostScenic
	} else {
		// We need to populate visible and then get the count.
		f.scenic = make([][]int, f.size)
		for i := range f.scenic {
			f.scenic[i] = make([]int, f.size)
		}

		for i := range f.scenic {
			for j := range f.scenic[i] {
				f.scenic[i][j] = f.ScenicScore(i, j)
			}
		}

		return f.calculateScenic()
	}
}

func (f *forest) ScenicScore(i, j int) int {
	upScore, downScore, rightScore, leftScore := 0, 0, 0, 0

	for y := i - 1; y >= 0; y-- {
		upScore++
		if f.trees[y][j] >= f.trees[i][j] {
			break
		}
	}

	if upScore == 0 {
		return 0
	}

	for y := i + 1; y <= f.size-1; y++ {
		downScore++
		if f.trees[y][j] >= f.trees[i][j] {
			break
		}
	}

	if downScore == 0 {
		return 0
	}

	for x := j - 1; x >= 0; x-- {
		leftScore++
		if f.trees[i][x] >= f.trees[i][j] {
			break
		}
	}

	if leftScore == 0 {
		return 0
	}

	for x := j + 1; x <= f.size-1; x++ {
		rightScore++
		if f.trees[i][x] >= f.trees[i][j] {
			break
		}
	}

	if rightScore == 0 {
		return 0
	}

	scenic := upScore * downScore * leftScore * rightScore
	return scenic
}

func (d *Day8) Solve() error {
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

func (d *Day8) Part1() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	var trees [][]int
	var grid *forest
	for index, line := range lines {
		if index == 0 {
			gridSize := len(line)
			trees = make([][]int, gridSize)
			for i := range trees {
				trees[i] = make([]int, gridSize)
			}
			grid = &forest{trees: trees, size: gridSize}
		}

		currRow := make([]int, grid.size)
		for ind, t := range line {
			num, err := strconv.Atoi(string(t))
			if err != nil {
				return nil, err
			}
			currRow[ind] = num
		}
		grid.PlaceTrees(currRow, index)
	}

	visibleCount := grid.calculateVisible()

	return visibleCount, nil
}

func (d *Day8) Part2() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	var trees [][]int
	var grid *forest
	var gridSize int
	for index, line := range lines {
		if index == 0 {
			gridSize = len(line)
			trees = make([][]int, gridSize)
			for i := range trees {
				trees[i] = make([]int, gridSize)
			}
			grid = &forest{trees: trees, size: gridSize}
		}

		currRow := make([]int, gridSize)
		for ind, t := range line {
			num, err := strconv.Atoi(string(t))
			if err != nil {
				return nil, err
			}
			currRow[ind] = num
		}
		grid.PlaceTrees(currRow, index)
	}

	scenicScore := grid.calculateScenic()

	return scenicScore, nil
}
