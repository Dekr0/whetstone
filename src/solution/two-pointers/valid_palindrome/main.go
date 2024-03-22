package main

/*
Less trivial. Since this method only look ahead by one letter, there's no 
guarantee that the rest of the given string is a palindrome. Thus, a skip count 
is here to keep track whether if it requires more than one skip.
    Case 1: only skip left => skip left, skip count + 1
    Case 2: only skip right => skip right, skip count + 1
    Case 3: either skip left or skip right => ignore this
        - there are two cases but we don't know at this point
            - one of the skip will produce a valid palindrome => delegate to 
            case 1 and case 2 
            - none of them will produce a valid palindrome => delegate to 
            case 4 in later iteration
        - try example "cucucu" and "cuaacu"
    Case 4: no skip => not possible to delete one letter to be a palindrome
*/
func solutionA(s string) bool {
    i := 0; j := len(s) - 1
    skipCount := 0
    for i <= j {
        if skipCount > 1 {
            return false
        }

        if s[i] != s[j] {
            skipCount++

            if s[i + 1] == s[j] && s[i] == s[j - 1] { // both skip is valid
                i++
                j--
                skipCount--
                continue
            } else if s[i + 1] == s[j] { // left skip is valid
                i++
                continue
            } else if s[i] == s[j - 1] { // right skip is valid
                j--
                continue
            } else { // both skip is not valid
                return false
            }
        }

        i++
        j--
    }

    return true
}

// Try both skip (function)
func solutionB(s string) bool {
    i := 0; j := len(s) - 1
    for i <= j {
        // induction => we don't need to whether if we skip before because 
        // we already valid the sub palindrome without one letter from either 
        // side. Thus, if boht of them return false, that's mean skipping once 
        // is not enought
        if s[i] != s[j] {
            return _validPalindrome(i + 1, j, s) || 
                _validPalindrome(i, j - 1, s)
        }
        i++
        j--
    }
    return true
}

func _validPalindrome(i int, j int, s string) bool {
    for i <= j {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
} 

// Try both skip (backtracking)
func solutionC(s string) bool {
    i := 0; j := len(s) - 1
    restore := make([]int, 2, 2)
    skipLeft := false; skipRight := true
    for i <= j {
        if s[i] == s[j] {
            i++
            j--
            continue
        } else if !skipLeft {
            restore[0] = i; restore[1] = j
            i++
            skipLeft = true
        } else if !skipRight {
            i = restore[0]; j = restore[1]
            j--
            skipRight = true
        } else {
            return false
        }
    }
    return true
}
