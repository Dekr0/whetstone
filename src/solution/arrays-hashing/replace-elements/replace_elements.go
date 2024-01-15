package replaceelements

func ReplaceElements(arr []int) []int {
    mx := arr[len(arr) - 1]
    arr[len(arr) - 1] = -1
    for i := len(arr) - 2; i >= 0; i-- {
        tmp := arr[i] 
        arr[i] = mx
        mx = max(tmp, mx)
    } 

    return arr
}
