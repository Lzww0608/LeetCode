func zigzagTraversal(grid [][]int) (ans []int) {
    m, n := len(grid), len(grid[0])

    cnt := 0
    for i := range m {
        if i & 1 == 0 {
            for j := range n {
                if cnt & 1 == 0 {
                    ans = append(ans, grid[i][j])
                }
                cnt++
            }
        } else {
            for j := n - 1; j >= 0; j-- {
                if cnt & 1 == 0 {
                    ans = append(ans, grid[i][j])
                }
                cnt++
            }
        }
    }

    return ans
}