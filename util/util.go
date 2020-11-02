package util

import (
    "bufio"
    "log"
    "os"
    "strconv"
    "strings"
)

/* IO methods */

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

/* Type Conversion Methods */

// StringToInt : converts a string value into an integer
func StringToInt(s string) int {
    res, err := strconv.Atoi(s)

    if err != nil {
        log.Fatal(err)
    }

    return res
}

// AlphabetToInt : converts a lowercase letter to its corresponding integer (from 0 to 25)
func AlphabetToInt(r rune) int {
    return int(r) - 97
}

// IntToAlphabet : converts an integer (0 to 25) into its corresponding alphabet
func IntToAlphabet(i int) rune {
    return rune(i + 97)
}

/* String Builder Methods */

// JoinStrings : concatenates input strings into one string
func JoinStrings(strs ...string) string {
    var sb strings.Builder

    for _, str := range strs {
        sb.WriteString(str)
    }

    return sb.String()
}

// JoinRunes : concatenates input runes into one string
func JoinRunes(runes ...rune) string {
    var sb strings.Builder

    for _, r := range runes {
        sb.WriteRune(r)
    }

    return sb.String()
}

/* CharCount Methods */

// CharCount : array representing the counts of each lowercase letter in a string
type CharCount struct {
    str string
    count [26]int
}

// NewCharCount : creates a new CharCount with the given string
func NewCharCount(str string) CharCount {
    var charCount CharCount
    charCount.str = str
    charCount.count = charCount.getCountFromString()
    return charCount
}

// GetString : returns the string represented by the charCount
func (cc CharCount) GetString() string {
    return cc.str
}

// GetCount : returns the count of the string represented by the charCount
func (cc CharCount) GetCount() [26]int {
    return cc.count
}

func (cc CharCount) getCountFromString() [26]int {
    var count [26]int

    for _, char := range cc.str {
        count[AlphabetToInt(char)]++
    }

    return count
}