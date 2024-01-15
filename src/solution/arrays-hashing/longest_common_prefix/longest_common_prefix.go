package longestcommonprefix

import (
	"strings"
)

func LongestCommonPrefix(strs []string) string {
    // The longest common prefix is limited by the length of the shortest string
    var prefix strings.Builder
    m_len := len(strs[0]); m_pos := 0
    for i := 1; i < len(strs); i++ {
        if len(strs[i]) < m_len {
            m_len = len(strs[i])
            m_pos = i
        }
    }
    for i := 0; i < m_len; i++ {
        for _, s := range strs {
            if s[i] != strs[m_pos][i] {
                return prefix.String()
            }
        }
        prefix.WriteByte(strs[m_pos][i])
    }
    return prefix.String()
}


func LongestCommonPrefixCryptic(strs[] string) string {
    var prefix strings.Builder
    for i := 0; i < len(strs[0]); i++ {
        for _, s := range strs {
            if len(s) == i || strs[0][i] != s[i] {
                return prefix.String()
            }
        }
        prefix.WriteByte(strs[0][i])
    }
    return prefix.String()
}
