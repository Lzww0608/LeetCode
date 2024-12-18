func minimumReplacement(nums []int) int64 {
    ans := 0
    n := len(nums)
    pre := nums[n-1]

    for i := n - 2; i >= 0; i-- {
        k := (nums[i] - 1) / pre 
        ans += k 
        pre = nums[i] / (k + 1)
    }

    return int64(ans)
}