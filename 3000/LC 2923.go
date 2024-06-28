func findChampion(grid [][]int) int {
    n := len(grid)
    ans := 0
    for i := 1; i < n; i++ {
        if grid[i][ans] == 1 {
            ans = i
        }
    }

    return ans
}
