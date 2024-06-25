func longestSubstring(s string, k int) int {
    n := len(s)
    if n < k {
        return 0
    }

    m := map[byte]int{}
    for i := range s {
        m[s[i]]++
    }

    i := 0
    for i < n && m[s[i]] >= k {
        i++
    }
    if i == n {
        return n
    }
    l := longestSubstring(s[:i], k)
    for i < n && m[s[i]] < k {
        i++
    }
    r := longestSubstring(s[i:], k)

    return max(l, r)
}
