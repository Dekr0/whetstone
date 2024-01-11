package main

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

var prime_encoding = [26]int{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,
61,67,71,73,79,83,89,97,101}

func IsAnagramAlt1(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    a := 0; b := 0

    for i := 0; i < len(s); i++ {
        a *= prime_encoding[s[i] - 'a']
        b *= prime_encoding[t[i] - 'a']
    }

    return a == b
}
