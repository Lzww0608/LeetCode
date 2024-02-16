var dir [4][2]int = [4][2]int{[2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}, [2]int{-1, 0},}
func spiralMatrixIII(rows int, cols int, rStart int, cStart int) (ans [][]int) {
    x, y := rStart, cStart
    n := rows * cols - 1
    ans = append(ans, []int{x, y})
    twd := 0
    for k := 1; n > 0; k++ {
        i, j := x, y
        for t := 0; t < k && n > 0; t++ {
            i, j = i + dir[twd][0], j + dir[twd][1]
            if i >= 0 && i < rows && j >= 0 && j < cols {
                ans = append(ans, []int{i, j})
                n--
            }
        }
        twd = (twd + 1) % 4;
        for t := 0; t < k && n > 0; t++ {
            i, j = i + dir[twd][0], j + dir[twd][1]
            if i >= 0 && i < rows && j >= 0 && j < cols {
                ans = append(ans, []int{i, j})
                n--
            }
        }
        x, y = i, j
        twd = (twd + 1) % 4;
        if n == 0 {
            return
        }
    }
    return
}