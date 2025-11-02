func deleteGreatestValue(grid [][]int) (ans int) {
    for i := range grid {
        sort.Ints(grid[i])
    }

    n := len(grid[0])
    for j := range n {
        cur := 0
        for i := range grid {
            cur = max(cur, grid[i][j])
        }

        ans += cur
    }

    return
}