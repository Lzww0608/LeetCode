func findNumberOfLIS(nums []int) (ans int) {
    n := len(nums)
    f := make([]int, n)
    cnt := make([]int, n)
    mx := 0
    for i, x := range nums {
        f[i], cnt[i] = 1, 1
        for j := i - 1; j >= 0; j-- {
            if x > nums[j] {
                if f[i] == f[j] + 1 {
                    cnt[i] += cnt[j]
                } else if f[i] < f[j] + 1 {
                    f[i] = f[j] + 1
                    cnt[i] = cnt[j]
                }
            }
        }
        mx = max(mx, f[i])
    }
    for i, x := range cnt {
        if mx == f[i] {
            ans += x
        }
    }
    return ans
}