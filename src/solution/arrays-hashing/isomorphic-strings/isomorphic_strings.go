package isomorphicstrings

func IsomorphicStrings(s string, t string) bool {
    transform := make(map[byte]byte)
    inverse := make(map[byte]byte)
    for i := 0; i < len(s); i++ {
        v, in := transform[s[i]]
        if in {
            if v != t[i] {
                return false
            }
        }
        v, in = inverse[t[i]]
        if in {
            if v != s[i] {
                return false
            }
        }
        transform[s[i]] = t[i]
        inverse[t[i]] = s[i]
    }
    return true
}

