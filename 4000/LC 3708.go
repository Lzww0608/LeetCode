func longestSubarray(nums []int) int {
    ans := 2 
    
    cur := 2
    for i := 2; i < len(nums); i++ {
        if nums[i] == nums[i - 1] + nums[i - 2] {
            cur++
        } else {
            cur = 2
        }

        ans = max(ans, cur)
    }

    return ans
}