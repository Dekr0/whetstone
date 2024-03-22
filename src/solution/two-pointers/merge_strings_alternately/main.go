package main

import "strings"

func mergeAlternately(word1 string, word2 string) string {
    var builder strings.Builder
    i := 0; j := 0
    runes1 := []rune(word1)
    runes2 := []rune(word2)
    for i < len(word1) && j < len(word2) {
        builder.WriteRune(runes1[i])
        i++
        builder.WriteRune(runes2[j])
        j++
    }
    if i != len(word1) {
        builder.WriteString(string(runes1[i:]))
    } else if j != len(word2) {
        builder.WriteString(string(runes2[j:]))
    }

    return builder.String()
}
