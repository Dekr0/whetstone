package main

import "math"

func IsAnagram(s string, t string) bool {
    /*
    1 <= s.length, t.length <= 5 * 104
    s and t consist of lowercase English letters (Important)
    */
    if len(s) != len(t) {
        return false
    }

    a := make([]int, 26, 26)
    b := make([]int, 26, 26)

    for i := 0; i < len(s); i++ {
        a[s[i] - 'a'] += 1
        b[t[i] - 'a'] += 1
    }

    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}

func UIntPow(x uint, y int) uint {
    if y == 0 {
        return 1
    }

    if x == 0 {
        return 0
    }

    if y == 1 {
        return x
    }

    r := uint(1)
    for i := 0; i < int(y); i++ {
        r *= x
    }

    return r
}

var prime_encoding = [26]uint{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,
61,67,71,73,79,83,89,97,101}

var sprime_encoding = [26]uint{2, 6, 15, 28, 55, 78, 119, 152, 207, 290, 341, 
444, 533, 602, 705, 848, 1003, 1098, 1273, 1420, 1533, 1738, 1909, 2136, 2425, 
2626}

var umod = UIntPow(2, 31) - 1 // magic number

var mod = math.Pow(2, 31) - 1

const cmod = 2147483647


func IsAnagramAlt1(s string, t string) bool {
    // slow because large floating point calculation is slow
    if len(s) != len(t) {
        return false
    }

    a := 1.0; b := 1.0

    for i := 0; i < len(s); i++ {
        a = math.Mod(a * float64(prime_encoding[s[i] - 'a']), mod)
        b = math.Mod(b * float64(prime_encoding[t[i] - 'a']), mod)
    }

    return a == b
}

func IsAnagramAlt2(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    a := uint(1); b := uint(1)

    for i := 0; i < len(s); i++ {
        a = a * prime_encoding[s[i] - 'a'] % umod
        b = b * prime_encoding[t[i] - 'a'] % umod
    }

    return a == b
}

func IsAnagramAlt3(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    a := uint(1); b := uint(1)

    for i := 0; i < len(s); i++ {
        a = a * prime_encoding[s[i] - 'a'] % cmod
        b = b * prime_encoding[t[i] - 'a'] % cmod
    }

    return a == b
}

func IsAnagramAlt4(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    a := uint(1); b := uint(1)

    for i := 0; i < len(s); i++ {
        a += sprime_encoding[s[i] - 'a']
        b += sprime_encoding[t[i] - 'a']
    }

    return a == b
}
