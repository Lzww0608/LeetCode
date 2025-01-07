func shareCandies(candies []int, k int) (ans int) {
    m := make(map[int]int)
    n := len(candies)
    for _, x := range candies {
        m[x]++
    }

    if k == 0 {
        return len(m)
    } else if k >= n {
        return 0
    }

    for l, r := 0, 0; r < n; r++ {
        if m[candies[r]]--; m[candies[r]] == 0 {
            delete(m, candies[r])
        }
        if r - l + 1 == k {
            ans = max(ans, len(m))
            m[candies[l]]++
            l++
        }
    }

    return
}