func minCosts(cost []int) []int {
    n := len(cost)
    ans := make([]int, n)
    mn := math.MaxInt32
    for i := range n {
        mn = min(mn, cost[i])
        ans[i] = mn
    }

    return ans
}