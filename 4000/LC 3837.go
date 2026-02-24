func delayedCount(nums []int, k int) []int {
    cnt := make(map[int]int)
    for _, x := range nums[k + 1:] {
        cnt[x]++
    }

    n := len(nums)
    ans := make([]int, n)
    for i, x := range nums[:n - k - 1] {
        ans[i] = cnt[x]
        cnt[nums[i + k + 1]]--
    }

    return ans
}