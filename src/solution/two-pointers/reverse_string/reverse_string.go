package main

func solution1(s []byte) {
    i := 0; j := len(s) - 1
    for i < j {
        tmp := s[j]
        s[j] = s[i]
        s[i] = tmp
        i++
        j--
    }
}

func solution2(s []byte) {
    for i := 0; i < len(s) / 2; i++ {
        s[i], s[len(s) - 1 - i] = s[len(s) - 1 - i], s[i]
    }
}
