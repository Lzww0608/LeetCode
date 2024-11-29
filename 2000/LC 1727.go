func largestSubmatrix(matrix [][]int) (ans int) {
    m, n := len(matrix), len(matrix[0])
    pre := make([][]int, n)
    for i := range pre {
        pre[i] = make([]int, m + 1)
    }

    for j := 0; j < n; j++ {
        for i := 0; i < m; i++ {
            if matrix[i][j] == 1 {
                pre[j][i+1] = pre[j][i] + 1
            } else {
                pre[j][i+1] = 0
            }
        }
    }

    for i := 0; i < m; i++ {
        a := make([]int, 0, n)
        for j := 0; j < n; j++ {
            a = append(a, pre[j][i+1])
        }
        sort.Slice(a, func(i, j int) bool {
            return a[i] > a[j]
        })
        cur := 0
        for j := range a {
            cur = max(cur, (j + 1) * a[j])
        }
        ans = max(ans, cur)
    }

    return
}