func equalPairs(grid [][]int) (ans int) {
    m := map[string]int{}
    n := len(grid)
    for _, x := range grid {
        m[sliceToString(x)]++
    }
    for j := 0; j < n; j++ {
        col := []int{}
        for i := 0; i < n; i++ {
            col = append(col, grid[i][j])
        }
        ans += m[sliceToString(col)]
    }
    return
}

func sliceToString(arr []int) string {
    var s string
    for _, x := range arr {
        s += strconv.Itoa(x) + ","
    }
    return s
}
