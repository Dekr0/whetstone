package main

import (
	"math/rand"
	"strings"
	"testing"
)

func CreateAnagram(l int) string {
    var builder strings.Builder
    for i := 0; i < l; i++ {
        builder.WriteByte(byte(97 + rand.Intn(26)))
    }

    return builder.String()
}

var s = CreateAnagram(10000)
var t = CreateAnagram(10000)

func TestIsAnagram(t *testing.T) {
    input := [][2]string {
        {"anagram", "nagaram"},
        {"rat", "car"},
    }

    output := []bool { true, false }

    for i := 0; i < len(input); i++ {
        actual := IsAnagram(input[i][0], input[i][1])
        if actual != output[i] {
            t.Fatalf(`Expected %t but get %t for test case %v`, output[i], 
            actual, input[i])
        }
    }
}

func TestIsAnagramAlt1(t *testing.T) {
    input := [][2]string {
        {"anagram", "nagaram"},
        {"rat", "car"},
    }

    output := []bool { true, false }

    for i := 0; i < len(input); i++ {
        actual := IsAnagramAlt1(input[i][0], input[i][1])
        if actual != output[i] {
            t.Fatalf(`Expected %t but get %t for test case %v`, output[i], 
            actual, input[i])
        }
    }
}

func TestIsAnagramAlt2(t *testing.T) {
    input := [][2]string {
        {"anagram", "nagaram"},
        {"rat", "car"},
    }

    output := []bool { true, false }

    for i := 0; i < len(input); i++ {
        actual := IsAnagramAlt2(input[i][0], input[i][1])
        if actual != output[i] {
            t.Fatalf(`Expected %t but get %t for test case %v`, output[i], 
            actual, input[i])
        }
    }
}

func TestIsAnagramAlt3(t *testing.T) {
    input := [][2]string {
        {"anagram", "nagaram"},
        {"rat", "car"},
    }

    output := []bool { true, false }

    for i := 0; i < len(input); i++ {
        actual := IsAnagramAlt3(input[i][0], input[i][1])
        if actual != output[i] {
            t.Fatalf(`Expected %t but get %t for test case %v`, output[i], 
            actual, input[i])
        }
    }
}

func BenchmarkIsAnagram(b *testing.B) {
    IsAnagram(s, t)
}

func BenchmarkIsAnagramAlt1(b *testing.B) {
    IsAnagramAlt1(s, t)
}

func BenchmarkIsAnagramAlt2(b *testing.B) {
    IsAnagramAlt2(s, t)
}

func BenchmarkIsAnagramAlt3(b *testing.B) {
    IsAnagramAlt3(s, t)
}
