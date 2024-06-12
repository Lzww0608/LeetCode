func productExceptSelf(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)
    ans[0] = 1

    for i := 0; i < n - 1; i++ {
        ans[i+1] = ans[i] * nums[i]
    }

    suf := 1
    for i := n - 2; i >= 0; i-- {
        suf *= nums[i+1]
        ans[i] *= suf
    }

    return ans

}
