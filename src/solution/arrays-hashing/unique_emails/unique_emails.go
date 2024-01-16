package uniqueemails

import (
	"fmt"
	"strings"
)

type Void struct {}

func NumUniqueEmails(emails []string) int {
    count := 0

    m := make(map[string]Void)

    var builder strings.Builder

    for _, e := range emails {
        split := strings.Split(e, "@")
        local, domain := strings.ReplaceAll(split[0], ".", ""), split[1]
        ignore_from := strings.IndexByte(local, '+')
        if ignore_from > -1 {
            local = local[:ignore_from]
        }
        builder.WriteString(local)
        builder.WriteByte('@')
        builder.WriteString(domain)
        pp := builder.String()
        builder.Reset()
        _, in := m[pp]
        if !in {
            m[pp] = Void{}
            count++
        }
    }


    return count
}

func Tokenizer(emails []string) int {
    count := 0
    var builder strings.Builder
    m := make(map[string]struct{})
    for _, email := range emails {
        domain := false
        for _, char := range email {
            switch char {
                case '+':
                    continue
                case '@':
                    domain = true
                case '.':
                    if !domain {
                        continue
                    }
                    fallthrough
                default:
                   builder.WriteRune(char) 
            }
        }
    }
    return count
}
