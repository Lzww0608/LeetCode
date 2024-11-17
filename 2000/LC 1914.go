func rotateGrid(grid [][]int, k int) [][]int {
    m, n := len(grid), len(grid[0])
    ans := make([][]int, m)
    for i := range ans {
        ans[i] = make([]int, n)
    }

    t := min(m, n) / 2
    for d := 0; d < t; d++ {
        tmp := make([]int, 0, (m + n - d * 4 - 2) * 2)
        for j := d; j < n - d; j++ {
            tmp = append(tmp, grid[d][j])
        }
        for i := d + 1; i < m - d; i++ {
            tmp = append(tmp, grid[i][n-1-d])
        }
        for j := n - d - 2; j >= d; j-- {
            tmp = append(tmp, grid[m-1-d][j])
        }
        for i := m - d - 2; i > d; i-- {
            tmp = append(tmp, grid[i][d])
        }

        shift := k % len(tmp)
        tmp = append(tmp[shift:], tmp[:shift]...)

        p := 0
        for j := d; j < n - d; j++ {
            ans[d][j] = tmp[p]
            p++
        }
        for i := d + 1; i < m - d; i++ {
            ans[i][n-1-d] = tmp[p]
            p++
        }
        for j := n - d - 2; j >= d; j-- {
            ans[m-1-d][j] = tmp[p]
            p++
        }
        for i := m - d - 2; i > d; i-- {
            ans[i][d] = tmp[p]
            p++
        }
    }
    
    return ans
}