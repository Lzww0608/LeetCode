func constructTransformedArray(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)
    for i, x := range nums {
        if x >= 0 {
            j := (i + x) % n 
            ans[i] = nums[j]
        } else {
            j := (i + x % n + n) % n 
            ans[i] = nums[j]
        }
    }

    return ans
}