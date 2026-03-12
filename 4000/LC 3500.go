func minimumCost(nums []int, cost []int, k int) int64 {
    n := len(nums)
    f := make([]int, n + 1)
    for i := range f {
        f[i] = math.MaxInt / 2
    }

    pre := make([]int, n + 1)
    for i := range cost {
        pre[i + 1] = pre[i] + cost[i] 
    }

    sum := 0
    f[0] = 0
    for i := 1; i <= n; i++ {
        sum += nums[i - 1]
        for j := range i {
            f[i] = min(f[i], f[j] + sum * (pre[i] - pre[j]) + k * (pre[n] - pre[j]))
        }
    }

    return int64(f[n])
}