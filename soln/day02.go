package soln

import (
    "fmt"
    "strings"

    "aoc2018/util"
)

func solve2() {
    data := util.ReadLines("data/day02.txt")
    counts := getCharCounts(data)

    checksum := calculateChecksum(counts)
    fmt.Printf("2A: %d\n", checksum)

    commonLetters := getCommonLetters(counts)
    fmt.Printf("2B: %s\n", commonLetters)
}

type charCounts []util.CharCount

func getCharCounts(ids []string) charCounts {
    counts := make(charCounts, len(ids))

    for i, id := range ids {
        counts[i] = util.NewCharCount(id)
    }

    return counts
}

func calculateChecksum(counts charCounts) int {
    countTwice, countThrice := 0, 0

    for _, charCount := range counts {
        if containsLetterNTimes(charCount, 2) {
            countTwice++
        }

        if containsLetterNTimes(charCount, 3) {
            countThrice++
        }
    }

    return countTwice * countThrice
}

func containsLetterNTimes(charCount util.CharCount, n int) bool {
    for _, count := range charCount.GetCount() {
        if count == n {
            return true
        }
    }

    return false
}

func getCommonLetters(counts charCounts) string {
    for i := 0; i < len(counts); i++ {
        for j := i + 1; j < len(counts); j++ {
            firstDiff, isCommonIds := getFirstDiffCharFromCommonIds(counts[i].GetString(), counts[j].GetString())

            if isCommonIds {
                return removeFirstRune(counts[i].GetString(), firstDiff)
            }
        }
    }

    return ""
}

func getFirstDiffCharFromCommonIds(idA string, idB string) (rune, bool) {
    if idA == idB {
        return '*', false
    }

    diffFound := false
    diff := '*'

    for i := 0; i < len(idA); i++ {
        if idA[i] != idB[i] {
            if diffFound {
                return '*', false
            }

            diffFound = true
            diff = rune(idA[i])
        }
    }

    return diff, diffFound
}

func removeFirstRune(str string, firstRune rune) string {
    var sb strings.Builder
    firstRuneFound := false

    for _, r := range str {
        if (!firstRuneFound && r == firstRune) {
            firstRuneFound = true
        } else {
            sb.WriteRune(r)
        }
    }

    return sb.String()
}