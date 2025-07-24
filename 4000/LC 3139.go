const MOD int = 1_000_000_007

func minCostToEqualizeArray(nums []int, cost1 int, cost2 int) int {
    n := len(nums)
    sum, mx, mn := 0, 0, math.MaxInt32
    for _, x := range nums {
        sum += x 
        mx = max(mx, x)
        mn = min(mn, x)
    }

    if cost1 * 2 <= cost2 || n <= 2 {
        return (n * mx - sum) % MOD * cost1 % MOD
    }

    solve := func(x int) int {
        d := x - mn 
        s := x * n - sum 
        if d * 2 <= s {
            return s / 2 * cost2 + (s & 1) * cost1
        }
        return (s - d) * cost2 + (d - (s - d)) * cost1 
    }

    pre := math.MaxInt
    for x := mx; ; x++ {
        cur := solve(x)
        if cur > pre {
            return pre % MOD
        }
        pre = cur
    }

    return -1
}