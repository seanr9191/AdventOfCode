package y2022

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/seanr9191/AdventOfCode/pkg/io/file"
)

type Day7 struct {
	Year      int
	Day       int
	InputFile string
	Logger    *zap.SugaredLogger
}

func (d *Day7) Solve() error {
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

type directory struct {
	parent   *directory
	children []*directory
	files    []*terminalFile
	name     string
	size     int
}

func (d *directory) FindChildren(name string) *directory {
	if d == nil {
		return nil
	}
	for _, child := range d.children {
		if strings.EqualFold(child.name, name) {
			return child
		}
	}
	return nil
}

type terminalFile struct {
	name string
	size int
}

func (d *Day7) Part1() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	directories := make(map[uuid.UUID]*directory)
	var currDirectory *directory
	for _, line := range lines {
		if strings.HasPrefix(line, "$ cd ") {
			currDirectoryName := strings.ReplaceAll(line, "$ cd ", "")
			if currDirectoryName != ".." {
				// Going down
				existingCurrDirectory := currDirectory.FindChildren(currDirectoryName)
				if existingCurrDirectory == nil {
					currDirectory = &directory{
						parent:   currDirectory,
						children: make([]*directory, 0),
						files:    make([]*terminalFile, 0),
						name:     currDirectoryName,
						size:     0,
					}
				} else {
					currDirectory = existingCurrDirectory
				}
				if currDirectory.parent != nil {
					currDirectory.parent.children = append(currDirectory.parent.children, currDirectory)
				}
				directories[uuid.New()] = currDirectory
			} else {
				// Going up
				// We also want to add the size of the directory we just left to the parent.
				currDirectory.parent.size += currDirectory.size

				// Make the move.
				currDirectory = currDirectory.parent
			}
		} else if strings.HasPrefix(line, "dir ") {
			// This is a child directory listing from the current directory
			// We don't actually handle anything here though.
			// Instead, we're assigning the child to the parent when we enter it.
		} else if unicode.IsDigit(rune(line[0])) {
			// This is a file within the current directory.
			parts := strings.Split(line, " ")
			size, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, err
			}

			// Let's add it as file
			currFile := &terminalFile{
				name: parts[1],
				size: size,
			}
			currDirectory.files = append(currDirectory.files, currFile)

			// And accumulate the size of the file in the directory.
			currDirectory.size += currFile.size
		}
	}

	// Move back to the head, adding the directory sizes to the parents.
	for currDirectory.parent != nil {
		currDirectory.parent.size += currDirectory.size
		// Make the move.
		currDirectory = currDirectory.parent
	}

	total := 0
	for _, dir := range directories {
		if dir.size <= 100000 {
			total += dir.size
		}
	}

	return total, nil
}

func (d *Day7) Part2() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	totalSpace := 70000000
	neededSpace := 30000000
	directories := make(map[uuid.UUID]*directory)
	var currDirectory *directory
	for _, line := range lines {
		if strings.HasPrefix(line, "$ cd ") {
			currDirectoryName := strings.ReplaceAll(line, "$ cd ", "")
			if currDirectoryName != ".." {
				// Going down
				existingCurrDirectory := currDirectory.FindChildren(currDirectoryName)
				if existingCurrDirectory == nil {
					currDirectory = &directory{
						parent:   currDirectory,
						children: make([]*directory, 0),
						files:    make([]*terminalFile, 0),
						name:     currDirectoryName,
						size:     0,
					}
				} else {
					currDirectory = existingCurrDirectory
				}
				if currDirectory.parent != nil {
					currDirectory.parent.children = append(currDirectory.parent.children, currDirectory)
				}
				directories[uuid.New()] = currDirectory
			} else {
				// Going up
				// We also want to add the size of the directory we just left to the parent.
				currDirectory.parent.size += currDirectory.size
				// Make the move.
				currDirectory = currDirectory.parent
			}
		} else if strings.HasPrefix(line, "dir ") {
			// This is a child directory listing from the current directory
			// We don't actually handle anything here though.
			// Instead, we're assigning the child to the parent when we enter it.
		} else if unicode.IsDigit(rune(line[0])) {
			// This is a file within the current directory.
			parts := strings.Split(line, " ")
			size, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, err
			}

			// Let's add it as file
			currFile := &terminalFile{
				name: parts[1],
				size: size,
			}
			currDirectory.files = append(currDirectory.files, currFile)

			// And accumulate the size of the file in the directory.
			currDirectory.size += currFile.size
		}
	}

	// Move back to the head, adding the directory sizes to the parents.
	for currDirectory.parent != nil {
		currDirectory.parent.size += currDirectory.size
		// Make the move.
		currDirectory = currDirectory.parent
	}

	currentFreeSpace := totalSpace - currDirectory.size
	minDeletion := neededSpace - currentFreeSpace

	var bestDir *directory
	for _, dir := range directories {
		if dir.size >= minDeletion && (bestDir == nil || dir.size < bestDir.size) {
			bestDir = dir
		}
	}

	return bestDir.size, nil
}
