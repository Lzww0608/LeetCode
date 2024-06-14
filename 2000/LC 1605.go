func restoreMatrix(rowSum []int, colSum []int) [][]int {
    m, n := len(rowSum), len(colSum)
    ans := make([][]int, m)
    for i := range ans {
        ans[i] = make([]int, n)
    }
    for i := range ans {
        x := rowSum[i]
        for j := range ans[i] {
            if colSum[j] <= x {
                ans[i][j] = colSum[j]
                x -= colSum[j]
                colSum[j] = 0
            } else {
                ans[i][j] = x
                colSum[j] -= x
                x = 0
                break
            }
        }
    }

    return ans
}
