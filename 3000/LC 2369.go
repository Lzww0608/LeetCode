func validPartition(nums []int) bool {
    n := len(nums)
    dp := make([]bool, n + 1)
    dp[0] = true

    check := func(i, k int) bool {
        if k == 2 {
            return nums[i] == nums[i+1]
        }
        if nums[i] == nums[i+2] && nums[i+1] == nums[i+2] {
            return true
        } else if nums[i] + 1 == nums[i+1] && nums[i+1] + 1 == nums[i+2] {
            return true
        }
        return false
    }

    for i := 1; i < n; i++ {
        if dp[i-1] == true {
            dp[i+1] = check(i - 1, 2) || dp[i+1]
        } 
        if i - 2 >= 0 && dp[i-2] == true {
            dp[i+1] = check(i - 2, 3) || dp[i+1]
        } 
    }
    return dp[n]
}
