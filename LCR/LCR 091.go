func minCost(costs [][]int) int {
    f := [3]int{}
    for _, cost := range costs {
        f[0], f[1], f[2] = min(f[1], f[2]) + cost[0], min(f[0], f[2]) + cost[1], min(f[0], f[1]) + cost[2]
    }

    return min(f[0], f[1], f[2])
}