func canPartitionGrid(grid [][]int) bool {
    m, n := len(grid), len(grid[0])
    sum := 0
    for _, v := range grid {
        for _, x := range v {
            sum += x
        }
    }

    cur := 0
    for _, v := range grid[:m - 1] {
        for _, x := range v {
            cur += x
        }
        if cur * 2 == sum {
            return true
        }
    }

    cur = 0
    for j := 0; j < n - 1; j++ {
        for i := 0; i < m; i++ {
            cur += grid[i][j]
        }
        if cur * 2 == sum {
            return true
        }
    }

    return false
}