func minimumTimeToInitialState(word string, k int) int {
    s := word + word
    n := len(word)
    z := make([]int, n)
    
    for i, l, r := 1, 1, 1; i < n; i++ {
        z[i] = max(min(z[i - l], r - i + 1), 0)
        for i + z[i] < n * 2 && s[i+z[i]] == s[z[i]] {
            l, r = i, i + z[i]
            z[i]++
        }

        if i % k == 0 && z[i] + i >= n {
            return (i + k - 1) / k
        }
    }
    return (n + k - 1) / k
}