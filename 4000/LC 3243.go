func shortestDistanceAfterQueries(n int, queries [][]int) []int {
    ans := make([]int, len(queries))
    f := make([]int, n)
    edge := make([][]int, n)
    for i := range f {
        f[i] = i
    }

    for k, q := range queries {
        l, r := q[0], q[1]
        edge[r] = append(edge[r], l)
        if f[l] + 1 < f[r] {
            f[r] = f[l] + 1
            for i := r + 1; i < n; i++ {
                f[i] = min(f[i], f[i-1] + 1)
                for _, j := range edge[i] {
                    f[i] = min(f[i], f[j] + 1)
                }
            }
        }

        ans[k] = f[n-1]
    }

    return ans
}