package nextgreaterelement

func pop(stack []int) ([]int, int, bool) {
    if len(stack) > 0 {
        return stack[:len(stack) - 1], stack[len(stack) - 1], true
    }

    return nil, -1, false
}

func peek(stack []int) (int, bool) {
    if len(stack) > 0 {
        return stack[len(stack) - 1], true
    }

    return -1, false
}

func NextGreaterElement(nums_1 []int, nums_2 []int) []int {
    i := len(nums_2) - 1

    m := make(map[int]int)
    mono_stack := make([]int, 0, len(nums_2)) 

    for ; i >= 0; i-- {
        for ; len(mono_stack) > 0; {
            top, _ := peek(mono_stack)
            if nums_2[i] <= nums_2[top] {
                m[nums_2[i]] = top
                mono_stack = append(mono_stack, i)
                break
            } else {
                mono_stack, _, _ = pop(mono_stack)
            }
        }
        if len(mono_stack) == 0 {
            m[nums_2[i]] = -1
            mono_stack = append(mono_stack, i)
        }
    }

    out := make([]int, 0, len(nums_1))

    for _, n := range nums_1 {
        if m[n] == -1 {
            out = append(out, -1)
        } else {
            out = append(out, nums_2[m[n]])
        }
    }

    return out
}
