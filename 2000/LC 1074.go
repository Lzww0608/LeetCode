func numSubmatrixSumTarget(matrix [][]int, target int) (ans int) {
    n, m := len(matrix), len(matrix[0])
    f := make([][]int, n + 1)
    for i := range f {
        f[i] = make([]int, m + 1)
    }

    for i := range matrix {
        for j, x := range matrix[i] {
            f[i+1][j+1] = f[i+1][j] + f[i][j+1] - f[i][j] + x
        }
    }

    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            cnt := map[int]int{0: 1}
            for k := 0; k < m; k++ {
                sum := f[j+1][k+1] - f[i][k+1]
                ans += cnt[sum - target]
                cnt[sum]++
            }
        }
    }

    return
}
