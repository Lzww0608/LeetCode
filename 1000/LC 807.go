func maxIncreaseKeepingSkyline(grid [][]int) (ans int) {
    n := len(grid)
    row := make([]int, n)
    col := make([]int, n)
    for i, t := range grid {
        row[i] = slices.Max(t)
        for j, x := range t {
            col[j] = max(col[j], x)
        }
    }
    
    for i, t := range grid {
        for j, x := range t {
            ans += min(col[j], row[i]) - x
        }
    }

    return
}