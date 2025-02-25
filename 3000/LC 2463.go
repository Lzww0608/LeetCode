func minimumTotalDistance(robot []int, factory [][]int) int64 {
    sort.Slice(factory, func(i, j int) bool {
        return factory[i][0] < factory[j][0]
    })
    sort.Ints(robot)
    n := len(robot)

    f := make([]int, n + 1)
    for i := range f {
        f[i] = math.MaxInt / 2
    }
    f[0] = 0

    for _, fa := range factory {
        for j := n; j > 0; j-- {
            cost := 0
            for k := 1; k <= min(j, fa[1]); k++ {
                cost += abs(robot[j-k] - fa[0])
                f[j] = min(f[j], f[j-k] + cost)
            }
        }
    }
    return int64(f[n])
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x 
}