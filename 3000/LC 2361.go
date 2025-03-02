func minimumCosts(regular []int, express []int, expressCost int) []int64 {
    a, b := 0, expressCost
    n := len(regular)
    ans := make([]int64, n)
    for i := 0; i < n; i++ {
        a, b = min(a, b) + regular[i], min(a + expressCost, b) + express[i]
        ans[i] = int64(min(a, b))
    }

    return ans
}