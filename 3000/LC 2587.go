func maxScore(nums []int) int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] > nums[j]
    })

    i, sum := 0, 0
    for i < len(nums) {
        sum += nums[i]
        if sum <= 0 {
            break
        }
        i++
    }

    return i 
}