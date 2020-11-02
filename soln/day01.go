package soln

import (
	"fmt"
	"aoc2018/util"
)

// Solve1 : prints the solutions for day 01
func Solve1() {
    data := util.ReadLines("data/day01.txt")
    resultingFrequency := calculateResultingFrequency(0, data)

    fmt.Printf("1A: %d\n", resultingFrequency)

    firstRepeatedFrequency := calculateFirstRepeatedFrequency(data)

    fmt.Printf("1B: %d\n", firstRepeatedFrequency)
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