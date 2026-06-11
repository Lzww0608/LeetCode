func maximumSum(nums []int, m int, l int, r int) int64 {
    n := len(nums)
    f := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
        if i == 0 {
            continue
        }
        for j := range f[i] {
            f[i][j] = math.MinInt / 2
        }
    }
    pre := make([]int, n + 1)
    for i, x := range nums {
        pre[i + 1] = pre[i] + x
    }

    ans := math.MinInt
    for i := 1; i <= m; i++ {
        q := []int{}
        for j := i * l; j <= n; j++ {
            k := j - l
            v := f[i - 1][k] - pre[k]
            for len(q) > 0 && f[i - 1][q[len(q) - 1]] - pre[q[len(q) - 1]] <= v {
                q = q[:len(q) - 1]
            }
            q = append(q, k)
            
            f[i][j] = max(f[i][j - 1], f[i - 1][q[0]] - pre[q[0]] + pre[j])

            if q[0] <= j - r {
                q = q[1:]
            }
        }

        ans = max(ans, f[i][n])
    }

    return int64(ans)
}