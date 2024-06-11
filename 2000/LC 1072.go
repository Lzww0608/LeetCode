func maxEqualRowsAfterFlips(matrix [][]int) (ans int) {
    m, n := len(matrix), len(matrix[0])
    mp := map[int]bool{}
    for i := range matrix {
        if mp[i] {
            continue
        } else if m - i <= ans {
            break
        }
        tmp := 1
        for k := i + 1; k < m; k++ {
            cnt := 0
            for j := range matrix[i] {
                if matrix[i][j] == matrix[k][j] {
                    cnt++
                } else {
                    cnt--
                }
            }
            if cnt == n || cnt == -n {
                mp[k] = true
                tmp++
            }  
        }
        ans = max(ans, tmp)
    }

    return
}