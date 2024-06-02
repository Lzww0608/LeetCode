func longestDecomposition(s string) (ans int) {
    n := len(s)
    l, r := 0, n
    for l <= r {
        k := 1
        for k = 1; k <= r - l; k++ {
            if s[l:l+k] == s[r-k:r] {
                if k ==  r - l {
                    ans++
                } else {
                    ans += 2
                }
                break
            }
        }
        l += k
        r -= k
    }

    return
}
