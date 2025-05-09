func findDiagonalOrder(mat [][]int) []int {
    m, n := len(mat), len(mat[0])
    ans := make([]int, 0, m * n)
    
    i, j := 0, 0 
    for k := 0; len(ans) < m * n; k++ {
        x, y := -1, 1 
        if k & 1 == 1 {
            x, y = 1, -1
        }

        ans = append(ans, mat[i][j])
        for i + x >= 0 && i + x < m && j + y >= 0 && j + y < n {
            i += x 
            j += y 
            ans = append(ans, mat[i][j])
        } 

        if k & 1 == 0 {
            if j + 1 < n {
                j += 1
            } else {
                i += 1  
            }
        } else {
            if i + 1 < m {
                i += 1
            } else {
                j += 1
            }
        }
    }

    return ans
}