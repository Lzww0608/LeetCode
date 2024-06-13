func longestAwesome(s string) (ans int) {
    //n := len(s)
    m := make([]int, 1 << 10)
    for i := range m {
        m[i] = math.MaxInt32
    }
    m[0] = 0
    mask := 0
    for i := range s {
        mask ^= 1 << int(s[i] - '0')
        if m[mask] == math.MaxInt32 {
            m[mask] = i + 1
        } else {
            ans = max(ans, i + 1 - m[mask])
        }
        for j := 0; j < 10; j++ {
            ans = max(ans, i + 1 - m[mask ^ (1 << j)])
        }
    }
    return
}
