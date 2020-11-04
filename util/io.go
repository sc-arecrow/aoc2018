package util

import (
	"bufio"
    "log"
    "os"
	"regexp"
	"strings"
)

// ReadLines : reads a file by line and returns and array containing each line as a string array split using regex as the delimiter
func ReadLines(fileName string, regexString string) [][]string {
    file, err := os.Open(fileName)

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    lines := [][]string{}
    scanner := bufio.NewScanner(file)
    re := regexp.MustCompile(regexString)

    for scanner.Scan() {
        rawText := scanner.Text()
        split := re.Split(rawText, -1)

        line := []string{}
        for i := range split {
            line = append(line, strings.TrimSpace(split[i]))
        }

        lines = append(lines, line)
    }

    return lines
}