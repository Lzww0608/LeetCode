const MOD int = 1_000_000_007
func solve(nums []int, queries [][]int) []int {
    n, m := len(nums), len(queries)
    ans := make([]int, m)

    sz := int(math.Sqrt(float64(m)))
    f := make([][]int, n)
    for i := range f {
        f[i] = make([]int, sz)
    }

    for y := 1; y < sz; y++ {
        for i := n - 1; i >= 0; i-- {
            cur := nums[i]
            if i + y < n {
                cur = (cur + f[i + y][y]) % MOD
            }
            f[i][y] = cur
        }
    }

    for i, v := range queries {
        x, y := v[0], v[1]
        if y >= sz {
            for j := x; j < n; j += y {
                ans[i] = (ans[i] + nums[j]) % MOD
            }
        } else {
            ans[i] = f[x][y]
        }
    }

    return ans
}