func maxValueOfCoins(piles [][]int, k int) int {
    n := len(piles)
    pre := make([][]int, n)
    for i, v := range piles {
        pre[i] = make([]int, len(v) + 1)
        for j, x := range v {
            pre[i][j+1] = pre[i][j] + x
        }
    }

    f := make([][]int, n + 1)
    f[0] = make([]int, k + 1)

    for i := 1; i <= n; i++ {
        f[i] = make([]int, k + 1)
        for j := 1; j <= k; j++ {
            sz := min(j, len(pre[i-1])-1)
            for t := 0; t <= sz; t++ {
                f[i][j] = max(f[i][j], f[i-1][j-t] + pre[i-1][t])
            } 
        }
    }
    
    return f[n][k]
}