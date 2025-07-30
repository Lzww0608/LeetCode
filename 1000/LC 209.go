func minSubArrayLen(target int, nums []int) int {
    ans := len(nums) + 1
    sum := 0
    for l, r := 0, 0; r < len(nums); r++ {
        sum += nums[r]
        for sum >= target {
            ans = min(ans, r - l + 1)
            sum -= nums[l]
            l++
        }
    }

    if ans > len(nums) {
        return 0
    }

    return ans
}