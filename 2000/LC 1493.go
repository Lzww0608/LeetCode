func longestSubarray(nums []int) int {
    ans, n := 0, len(nums)
    l, r, cnt := 0, 0, 0
    for r < n {
        if nums[r] == 0 {
            cnt++
            for cnt > 1 {
                if nums[l] == 0 {
                    cnt--
                }
                l++
            }
        }
        ans = max(ans, r - l)
        r++
    }
    return ans
}
