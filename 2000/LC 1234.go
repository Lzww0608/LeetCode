func balancedString(s string) int {
    n := len(s)
    m := map[byte]int {}
    for i := range s {
        m[s[i]]++
    }

    if m['W'] == n / 4 && m['Q'] == n / 4 && m['E'] == n / 4 && m['R'] == n / 4 {
        return 0
    }

    ans, l := n, 0    
    for r := range s {
        m[s[r]]--
        for m['W'] <= n / 4 && m['Q'] <= n / 4 && m['E'] <= n / 4 && m['R'] <= n / 4 {
            ans = min(ans, r - l + 1)
            m[s[l]]++
            l++
        }
    }

    return ans

}
