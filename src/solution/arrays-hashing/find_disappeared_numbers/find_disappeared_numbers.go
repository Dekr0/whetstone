package finddisappearednumbers

func FindDisappearedNumbers(nums []int) {
    disappeared := make([]int, len(nums), len(nums))
    for i := 0; i < len(disappeared); i++ {
        disappeared[i] = i + 1
    }
}
