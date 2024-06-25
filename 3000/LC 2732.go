func goodSubsetofBinaryMatrix(grid [][]int) []int {
    n := len(grid[0])
    m := make([]int, 1 << n)
    for i := range m {
        m[i] = -1
    }

    for i := range grid {
        mask := 0
        for j, x := range grid[i] {
            mask |= x << j
        }
        if mask == 0 {
            return []int{i}
        }
        m[mask] = i
    }

    u := (1 << n) - 1
    for x := 1; x < 1 << n; x++ {
        i := m[x]
        if i < 0 {
            continue
        }
        c := u ^ x
        for y := c; y != 0; y = (y - 1) & c {
            j := m[y]
            if j >= 0 {
                return []int{min(i, j), max(i, j)}
            }
        }
    }

    return nil
}