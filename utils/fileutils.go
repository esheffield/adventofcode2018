package utils

import (
	"bufio"
	"os"
)

func ReadLines(path string) ([]string, error) {
	var lines []string
	file, err := os.Open(path)

	if err != nil {
		return lines, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		lines = append(lines, str)
	}

	return lines, nil
}
