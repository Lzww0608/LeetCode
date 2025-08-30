func maximumMatchingIndices(a []int, b []int) (ans int) {
    n := len(a)
    for i := 0; i < n; i++ {
        cur := 0
        for j := 0; j < n; j++ {
            if a[j] == b[(i + j) % n] {
                cur++
            }
        }
        ans = max(ans, cur)
    }

    return 
}