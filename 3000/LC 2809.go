func minimumTime(a []int, b []int, x int) int {
    n := len(a)
    id := make([]int, n)
    for i := range id {
        id[i] = i
    }

    sort.Slice(id, func(i, j int) bool {
        return b[id[i]] < b[id[j]]
    })

    f := make([]int, n + 1)
    sum1, sum2 := 0, 0
    for i := 0; i < n; i++ {
        sum1 += a[i]
        sum2 += b[i]

        for j := i + 1; j > 0; j-- {
            f[j] = max(f[j], f[j-1] + a[id[i]] + b[id[i]] * j)
        }
    }

    for t := 0; t <= n; t++ {
        if sum1 + sum2 * t - f[t] <= x {
            return t
        }
    }

    return -1
}