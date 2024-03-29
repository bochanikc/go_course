package datafile

import (
	"bufio"
	"os"
)

func GetStrings(filename string) ([]string, error) {
	var lines []string
	var line string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line = scanner.Text()
		if err != nil {
			return nil, err
		}
		lines = append(lines, line)
	}

	err = file.Close()

	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return lines, scanner.Err()
	}
	return lines, nil
}
