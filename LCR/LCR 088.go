func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    a, b := cost[0], cost[1]

    for i := 2; i < n; i++ {
        b, a = min(a, b) + cost[i], b
    }

    return min(a, b)
}