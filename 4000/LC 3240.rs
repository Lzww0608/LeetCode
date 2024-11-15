func minFlips(grid [][]int) (ans int) {
    m, n := len(grid), len(grid[0])
    one := 0
    cnt := 0
    for i := 0; i < (m + 1) / 2; i++ {
        for j := 0; j < (n + 1) / 2; j++ {
            if (m & 1 == 1 && i + 1 == (m + 1) / 2) && (n & 1 == 1 && j + 1 == (n + 1) / 2) {
                if grid[i][j] == 1 {
                    ans++
                }
            } else if m & 1 == 1 && i + 1 == (m + 1) / 2 {
                k := n - 1 - j 
                if grid[i][j] == 1 {
                    one++
                } 
                if grid[i][k] == 1 {
                    one++
                } 
                if grid[i][j] != grid[i][k] {
                    cnt++
                }
            } else if n & 1 == 1 && j + 1 == (n + 1) / 2 {
                k := m - 1 - i 
                if grid[i][j] == 1 {
                    one++
                } 
                if grid[k][j] == 1 {
                    one++
                } 
                if grid[i][j] != grid[k][j] {
                    cnt++
                }
            } else {
                x := grid[i][j] + grid[m-1-i][j] + grid[i][n-1-j] + grid[m-1-i][n-1-j]
                ans += min(x, 4 - x)
            }
        }
    }

    return ans + max(cnt, min(one % 4, 4 - one % 4))
}