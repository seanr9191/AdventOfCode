package file

import (
	"bufio"
	"os"
)

type Reader struct {
	Path string
}

func (fr *Reader) GetContents() ([]string, error) {
	file, err := os.Open(fr.Path)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return lines, err
}
