func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    m, n := len(g), len(s)

    ans := 0
    for i, j := 0, 0; i < m && j < n; {
        if g[i] <= s[j] {
            i++
            j++
            ans++
        } else {
            j++
        }
    }

    return ans
}
