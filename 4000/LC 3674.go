func minOperations(nums []int) int {
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[0] {
            return 1
        }
    }

    return 0
}