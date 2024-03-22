package main

func abs(n int) int {
    if (n < 0) {
        return -n
    } else {
        return n
    }
}

func containNearByDuplicate(nums []int, k int) bool {
    if k == 0 {
        return false
    }
    
    /* Implicit Nested Loop because i is increment by one and j is reset under
    specific condition. The process of searching for duplicate is an O(N) 
    operation. */
    /* A single loop does not guarantee a linear runtime if index is being 
    reset. It's unwrapping a nested loop into a single loop */
    /* If hashing helps the ease of elminate the needs of condition check and 
    reduce runtime down to linear, use it */
    /* It's common to augment different methods to solve a problem */
    /* Walking every element over through the sliding window is not 
    necessarliy a must when using sliding window method. It can take the form 
    of another DA (e.g. set). If the sliding window is sufficient large, 
    linearly walking will induce performance pentaly */
    for i, j := 0, 1; i < len(nums) - 1; {
        less := abs(i - j) <= k
        same := j < len(nums) && nums[i] == nums[j]
        if same && less {
            return true
        }
        if less {
            j++
        } else {
            i++
            j = i + 1
        }
    }

    return false
}

func solution(nums []int, k int) bool {
    l := 0
    /* memory match up to the input since it does not remove anything */
    set := make(map[int]struct{})
    for r := 0; r < len(nums); {
        if r - l > k {
            delete(set, nums[l])
            l++
            continue
        }
        _, in := set[nums[r]]
        if in {
            return true
        }
        r++
    }
    return false
}

func final(nums []int, k int) bool {
    // 1,2,3,1, k = 2
    set := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        if idx, in := set[nums[i]]; in {
            if abs(idx - i) <= k {
                return true
            }
        } else {
            set[nums[i]] = i
        }
    }
    return false
}

func main() {
    containNearByDuplicate([]int{1,2,3,4,5,6,7,8,9,9}, 3)
}
