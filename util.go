package main

import (
	"bufio"
	"log"
	"os"
)

// ReadLines : reads a file by line and returns lines in a string array
func ReadLines(fileName string) []string {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
