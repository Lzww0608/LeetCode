func maxProduct(words []string) (ans int) {
    n := len(words)
    m := make([]int, n)
    for i, s := range words {
        mask := 0
        for _, c := range s {
            mask |= 1 << (c - 'a')
        }
        m[i] = mask
    }

    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if m[i] & m[j] == 0 {
                ans = max(ans, len(words[i]) * len(words[j]))
            }
        }
    }

    return
}