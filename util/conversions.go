package util

import (
    "log"
    "strconv"
)

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
