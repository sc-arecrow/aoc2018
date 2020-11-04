package soln

import (
    "fmt"
    
	"aoc2018/util"
)

func solve1() {
    data := util.ReadLines("data/day01.txt", " ")
    frequencies := getFrequencies(data)

    resultingFrequency := calculateResultingFrequency(0, frequencies)
    fmt.Printf("1A: %d\n", resultingFrequency)

    firstRepeatedFrequency := calculateFirstRepeatedFrequency(frequencies)
    fmt.Printf("1B: %d\n", firstRepeatedFrequency)
}

func getFrequencies(data [][]string) []string {
    frequencies := []string{}

    for _, datum := range data {
        frequencies = append(frequencies, datum[0])
    }

    return frequencies
}

func calculateResultingFrequency(initialFrequency int, frequencies []string) int {
    for _, i := range frequencies {
        frequency := util.StringToInt(i)
        initialFrequency += frequency
    }

    return initialFrequency
}

func calculateFirstRepeatedFrequency(frequencies []string) int {
    set := make(map[int]int)
    currentFrequency := 0
    repeatedFrequencyFound := false

    for !repeatedFrequencyFound {
        for _, i := range frequencies {
            frequency := util.StringToInt(i)
            currentFrequency += frequency
            _, containsKey := set[currentFrequency]
    
            if (containsKey) {
                repeatedFrequencyFound = true
                break
            } else {
                set[currentFrequency] = 1
            }
        }
    }
    
    return currentFrequency
}