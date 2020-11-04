package util

import (
	"strings"
)

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