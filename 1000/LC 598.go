func maxCount(m int, n int, ops [][]int) int {
    x, y := m, n
    for _, op := range ops {
        x = min(x, op[0])
        y = min(y, op[1])
    }

    return x * y
}
