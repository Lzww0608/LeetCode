const MOD int = 1_000_000_007
func countPartitions(nums []int, k int) int {
    mx := []int{}
    mn := []int{}

    cur, n, l := 0, len(nums), 0
    f := make([]int, n + 1)
    f[0] = 1
    for i, x := range nums {
        for len(mx) > 0 && nums[mx[len(mx) - 1]] <= x {
            mx = mx[:len(mx) - 1]
        }
        mx = append(mx, i)
        for len(mn) > 0 && nums[mn[len(mn) - 1]] >= x {
            mn = mn[:len(mn) - 1]
        }
        mn = append(mn, i)

        cur = (cur + f[i]) % MOD
        for len(mn) > 0 && len(mx) > 0 && nums[mx[0]] - nums[mn[0]] > k {
            if l == mx[0] {
                mx = mx[1:]
            }
            if l == mn[0] {
                mn = mn[1:]
            }

            cur = (cur - f[l] + MOD) % MOD
            l++
        }

        cur %= MOD
        f[i + 1] = cur
    }

    return f[n]
}