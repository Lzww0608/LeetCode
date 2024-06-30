func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    for i := 1; i < n; i++ {
        triangle[i][i] += triangle[i-1][i-1]
        for j := i - 1; j > 0; j-- {
            triangle[i][j] += min(triangle[i-1][j], triangle[i-1][j-1])
        }
        triangle[i][0] += triangle[i-1][0]
    }
    

    return slices.Min(triangle[n-1])
}
