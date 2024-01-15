package lengthoflastword

import "testing"

func TestLengthOfLastWord(t *testing.T) {
    input := []string{
        "Hello World",
        "   fly me   to   the moon  ",
        "luffy is still joyboy",
        "a",
    }

    output := []int{5, 4, 6, 1}
    for i := 0; i < len(input); i++ {
        get := LengthOfLastWord(input[i])
        if get != output[i] {
            t.Fatalf("Test Case %v failed with %v", input[i], get)
        }
    }
}
