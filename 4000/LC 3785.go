func minSwaps(nums []int, forbidden []int) int {
    n := len(nums)
    cnt := make(map[int]int)
    for i := range n {
        cnt[nums[i]]++
        cnt[forbidden[i]]++
        if cnt[nums[i]] > n || cnt[forbidden[i]] > n {
            return -1
        }
    }

    clear(cnt)
    sum, mx := 0, 0
    for i := range n {
        if nums[i] == forbidden[i] {
            sum++
            cnt[nums[i]]++
            mx = max(mx, cnt[nums[i]])
        }
    }

    return max(mx, (sum + 1) / 2)
}