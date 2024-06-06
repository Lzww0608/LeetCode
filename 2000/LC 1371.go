func findTheLongestSubstring(s string) (ans int) {
    state := 0
    a := "aeiou"
    pos := make([]int, 1 << 5)
    for i := range pos {
        pos[i] = 0x3f3f3f3f
    }
    pos[0] = -1
    for i := range s {
        for k := 0; k < 5; k++ {
            if s[i] == a[k] {
                state ^= (1 << k)
            }
        }
        pos[state] = min(pos[state], i)
        ans = max(ans, i - pos[state])
    }

    return
}
