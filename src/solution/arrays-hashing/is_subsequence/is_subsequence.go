package issubsequence

import "strings"

func IsSubsequence(s string, t string) bool {
    if len(t) == 0 && len(s) == 0 {
        return true
    }

    if len(t) == 0 {
        return false
    }

    if len(s) == 0 {
        return true
    }

    if len(s) == 1 {
        return strings.Contains(t, s)
    }

    target := 0
    target_pos := -1
    for i := 0; i < len(t); i++ {
        if target_pos < 0 {
            if s[target] == t[i] {
                target_pos = i
            }
        } else if target + 1 < len(s) {
            if s[target + 1] == t[i] {
                if target_pos < i {
                    target_pos = i
                    target += 1
                }
            }
        }
    }

    return target == len(s) - 1
}


func IsSubsequenceAlt(s string, t string) bool {
    // it's fairly simple kekw
    // lesson, start with minimal state, pen and pencil
    k := 0
    for i := 0; i < len(t); i++ {
        if k < len(s) && s[k] == t[i] {
            k++
        }
    }

    return k == len(s)
}
