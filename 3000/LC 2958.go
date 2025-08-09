func maxSubarrayLength(nums []int, k int) (ans int) {
    cnt := make(map[int]int)
    n := len(nums)

    for l, r := 0, 0; r < n; r++{
        cnt[nums[r]]++

        for cnt[nums[r]] > k {
            cnt[nums[l]]--
            l++
        }

        ans = max(ans, r - l + 1)
    }

    return
}