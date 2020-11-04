package soln

import (
    "fmt"
    "strings"

    "aoc2018/util"
)

func solve2() {
    data := util.ReadLines("data/day02.txt", " ")
    ids := getIds(data)
    counts := getCharCountsArray(ids)

    checksum := calculateChecksum(counts)
    fmt.Printf("2A: %d\n", checksum)

    commonLetters := getCommonLetters(counts)
    fmt.Printf("2B: %s\n", commonLetters)
}

func getIds(data [][]string) []string {
    ids := []string{}

    for _, datum := range data {
        ids = append(ids, datum[0])
    }

    return ids
}

type charCountsArray []util.CharCount

func getCharCountsArray(ids []string) charCountsArray {
    counts := make(charCountsArray, len(ids))

    for i, id := range ids {
        counts[i] = util.NewCharCount(id)
    }

    return counts
}

func calculateChecksum(charCounts charCountsArray) int {
    countTwice, countThrice := 0, 0

    for _, charCount := range charCounts {
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

func getCommonLetters(charCounts charCountsArray) string {
    for i := 0; i < len(charCounts); i++ {
        for j := i + 1; j < len(charCounts); j++ {
            firstDiff, isCommonIds := getFirstDiffCharFromCommonIds(charCounts[i].GetString(), charCounts[j].GetString())

            if isCommonIds {
                return removeFirstRune(charCounts[i].GetString(), firstDiff)
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