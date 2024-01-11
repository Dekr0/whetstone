package main

import (
	"testing"
)

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
