func spiralArray(array [][]int) (ans []int) {
    if len(array) == 0 {
        return 
    }
    m, n := len(array), len(array[0])
    i, j := 0, 0

    for k := 0; len(ans) < m * n; k++ {
        i, j = k, k
        for ; len(ans) < m * n && j < n - k; j++ {
            ans = append(ans, array[i][j])
        }
        for i, j = i + 1, j - 1; len(ans) < m * n && i < m - k; i++ {
            ans = append(ans, array[i][j])
        }
        for i, j = i - 1, j - 1; len(ans) < m * n && j >= k; j-- {
            ans = append(ans, array[i][j])
        }
        for i, j = i - 1, j + 1; len(ans) < m * n && i > k; i-- {
            ans = append(ans, array[i][j])
        }
    }

    return 
}