func paintWalls(cost []int, time []int) int {
    n := len(cost)
    f := make([]int, n + 1)
    for i := range f {
        f[i] = math.MaxInt / 2
    }
    f[0] = 0

    for i, x := range cost {
        t := time[i] + 1
        for j := n; j >= 0; j-- {
            f[j] = min(f[j], f[max(j - t, 0)] + x)
        }
    }
    
    return f[n]
}