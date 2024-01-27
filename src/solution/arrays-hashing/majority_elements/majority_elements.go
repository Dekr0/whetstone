package majorityelements

func MajorityElements(nums []int) int {
    // This algorithm solely relieve the fact that majority elements will 
    // always exist in the list of elements. Majority => more than half of the
    // the size of elements

    // Boyer-Moore Majority Voting Algorithm
    votes := 0
    candidate := -1
    i := 0
    for ; i < len(nums); i++ {
        if votes == 0 {
            candidate = nums[i]
            votes = 1
        } else if candidate == nums[i] {
            votes++
        } else {
            votes--
        }
    }
    /* skip checking process in Boyer-Moore Majority Voting Algorithm
    count := 0
    for i = 0; i < len(nums); i++ {
        if (nums[i] == candidate) {
            count++
        }
    }
    if count > len(nums) / 2 {
        return candidate
    } else {
        return -1
    }*/

    return candidate
}
