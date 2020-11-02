package util

import (
    "bufio"
    "log"
    "os"
    "strconv"
)

// ReadLines : reads a file by line and returns lines in a string array
func ReadLines(fileName string) []string {
    file, err := os.Open(fileName)

    if (err != nil) {
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

// StringToInt : converts a string value into an integer
func StringToInt(s string) int {
    res, err := strconv.Atoi(s)

    if (err != nil) {
        log.Fatal(err)
    }

    return res
}