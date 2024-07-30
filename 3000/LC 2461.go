func maximumSubarraySum(nums []int, k int) int64 {
    ans := 0
    m := map[int]int{}
    l, n := 0, len(nums)
    cur := 0
    for r := 0; r < n; r++ {
        m[nums[r]]++
        cur += nums[r]
        for r - l >= k {
            m[nums[l]]--
            cur -= nums[l]
            if m[nums[l]] == 0 {
                delete(m, nums[l])
            }
            l++
        }

        if r - l + 1 == k && len(m) == k {
            ans = max(ans, cur)
        }
    }

    return int64(ans)
}