package utils

import (
	"bufio"
	"os"
)

// ReadFile Read entire file into an array of string
func ReadFile(path string) ([]string, error) {
	var lines []string
	file, err := os.Open(path)

	if err != nil {
		return lines, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return lines, err
	}

	return lines, nil
}
