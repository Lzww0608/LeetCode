func minPenalty(period int, a []int, b []int) (ans int) {
    mx := slices.Max(a)

    for _, x := range b {
        x %= period
        if x >= mx {
            ans = max(ans, period - x)
        }
        
    }

    return
}