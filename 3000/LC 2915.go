func lengthOfLongestSubsequence(nums []int, target int) int {
    f := make([]int, target + 1)
    f[0] = 1
    for _, x := range nums {
        for i := target; i >= x; i-- {
            if f[i-x] != 0 {
                f[i] = max(f[i], f[i-x] + 1)
            }
            
        }
    }

    return f[target] - 1
}