func minTotalDistance(grid [][]int) (ans int) {
    x, y := []int{}, []int{}
    for i := range grid {
        for j, t := range grid[i] {
            if t == 1 {
                x = append(x, i)
                y = append(y, j)
            }
        }
    }

    sort.Ints(x)
    sort.Ints(y)
    n := len(x) / 2
    for i := range x {
        ans += abs(x[i] - x[n]) + abs(y[i] - y[n])
    }
    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}