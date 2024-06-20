func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
    ans := 0
    for _, c := range composition {
        l, r := 0, int(1e11)
        for l <= r {
            mid := l + ((r - l) >> 1)
            t := 0
            for i := 0; i < n; i++ {
                t += max(0, c[i] * mid - stock[i]) * cost[i] 
            }
            if t <= budget {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
        ans = max(ans, r)
    }
    return ans
}
