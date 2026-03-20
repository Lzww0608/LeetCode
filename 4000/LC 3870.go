func countCommas(n int) (ans int) {
    check := func(x int) int {
        res := 0
        for x >= 1000 {
            x /= 1000
            ans++
        }
        return res
    }

    for i := 1; i <= n; i++ {
        ans += check(i)
    }

    return
}