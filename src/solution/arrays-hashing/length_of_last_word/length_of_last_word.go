package lengthoflastword

func LengthOfLastWord(s string) int {
    j := -1
    i := len(s) - 1
    for ; i >= 0; i-- {
        if s[i] != ' ' {
            j = i
            break
        }
    }
    if j == -1 {
        return 0
    }
    for i = j; i >= 0; i-- {
        if s[i] == ' ' {
            break;
        }
    }
    return j - i
}
