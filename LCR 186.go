func checkDynasty(places []int) bool {
    mx, mn := math.MinInt32, math.MaxInt32
    m := map[int]int{}
    for _, x := range places {
        if x == 0 {
            continue
        }
        m[x]++
        if m[x] > 1 {
            return false
        }
        mx = max(mx, x)
        mn = min(mn, x)
    }
    if mx - mn <= 4 {
        return true
    }
    return false
}