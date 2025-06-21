func sumCounts(nums []int) (ans int) {
    n := len(nums)
    for i := 0; i < n; i++ {
        m := make(map[int]struct{})
        for j := i; j < n; j++ {
            m[nums[j]] = struct{}{}
            ans += len(m) * len(m)
        }
    }

    return
}