func countCornerRectangles(a [][]int) (ans int) {
    m, n := len(a), len(a[0])
    for i := 0; i < m; i++ {    
        for j := i + 1; j < m; j++ {
            cnt := 0
            for k := 0; k < n; k++ {
                if a[i][k] + a[j][k] == 2 {
                    cnt++
                }
            }
            ans += cnt * (cnt - 1) / 2
        }
    }

    return
}