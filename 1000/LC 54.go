func spiralOrder(matrix [][]int) (ans []int) {
    m, n := len(matrix), len(matrix[0])
    for k := 0; len(ans) < m * n; k++ {
        i, j := k, k
        for len(ans) < m * n && j < n - k {
            ans = append(ans, matrix[i][j])
            j++
        }
        j--
        i++
        for len(ans) < m * n && i < m - k {
            ans = append(ans, matrix[i][j])
            i++
        }
        i--
        j--
        for len(ans) < m * n && j >= k {
            ans = append(ans, matrix[i][j])
            j--
        }
        j++
        i--
        for len(ans) < m * n && i > k {
            ans = append(ans, matrix[i][j])
            i--
        }
    }

    return ans
}
