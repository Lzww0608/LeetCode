func splitArraySameAverage(nums []int) bool {
    if len(nums) == 1 {
        return false
    }
    sum := 0
    for _, x := range nums {
        sum += x 
    }
    if sum == 0 && len(nums) > 1 {
        return true
    }
    

    dp := make([]int, sum + 1)
    dp[0] = 1

    for _, x := range nums {
        for j := sum / 2; j >= x; j-- {
            dp[j] |= dp[j - x] << 1
            if j != 0 && j * len(nums) % sum == 0 && (1 << (j *len(nums) / sum)) & dp[j] > 0 {
                return true
            }
        }
    }

    return false
}