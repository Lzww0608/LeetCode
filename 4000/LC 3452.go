func sumOfGoodNumbers(nums []int, k int) (ans int) {
    n := len(nums)
    for i := range nums {
        if (i - k < 0 || nums[i] > nums[i - k]) && (i + k >= n || nums[i] > nums[i + k]) {
            ans += nums[i]
        }
    }

    return
}