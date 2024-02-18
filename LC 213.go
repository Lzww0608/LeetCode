func rob(nums []int) int {
    return max(nums[0], rob1(nums[1:]), rob1(nums[0:len(nums)-1]))
}

func rob1(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    a, b := 0, nums[0]
    ans := b
    for i := 1; i < n; i++ {
        a, b = b, max(b, a + nums[i])
        ans = max(ans, b)
    }
    return ans
}