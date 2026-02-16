func dominantIndices(nums []int) (ans int) {
    n := len(nums)
    sum := nums[n - 1]

    for i := n - 2; i >= 0; i-- {
        if nums[i] > sum / (n - i - 1) {
            ans++
        }
        sum += nums[i]
    }

    return
}