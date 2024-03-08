func kSum(nums []int, k int) int64 {
    sum, total := 0, 0
    n := len(nums)
    for i, x := range nums {
        if x > 0 {
            sum += x
            total += x
        } else {
            total -= x
            nums[i] = -x
        }
    }
    slices.Sort(nums)

    k_largest := sort.Search(total, func(limit int) bool {
        cnt := 1
        var dfs func(int, int)
        dfs = func(i, s int) {
            if cnt == k || i == n || s + nums[i] > limit {
                return
            }
            cnt++
            dfs(i + 1, s + nums[i])
            dfs(i + 1, s)
        }
        dfs(0, 0)
        return cnt == k
    })
    return int64(sum - k_largest)
}