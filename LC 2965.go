func findMissingAndRepeatedValues(grid [][]int) []int {
    n := len(grid)
    m := n * n
    d1, d2 := m * (m + 1) * (2 * m + 1) / 6, (m + 1) * m / 2
    for i := range grid {
        for _, x := range grid[i] {
            d1 -= x * x
            d2 -= x
        }
    }
    d := d1 / d2

    return []int{(d - d2) / 2, (d + d2) / 2}
}