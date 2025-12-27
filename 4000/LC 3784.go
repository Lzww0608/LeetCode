const N int = 26
func minCost(s string, cost []int) int64 {
    m := [N]int{}
    sum := 0
    for i := range cost {
        x := int(s[i] - 'a')
        sum += cost[i]
        m[x] += cost[i]
    }

    ans := sum
    for _, x := range m {
        if x > 0 {
            ans = min(ans, sum - x)
        }
    }

    return int64(ans)
}