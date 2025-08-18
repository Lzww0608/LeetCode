func maxProduct(words []string) (ans int) {
    n := len(words)
    mask := make([]int, n)
    for i, word := range words {
        for _, c := range word {
            mask[i] |= 1 << int(c - 'a')
        }
    }

    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if mask[i] & mask[j] == 0 {
                ans = max(ans, len(words[i]) * len(words[j]))
            }
        }
    }

    return
}