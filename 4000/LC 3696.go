func maxDistance(words []string) int {
    n := len(words)
    a, b := 0, 0

    for i := n - 1; i > 0; i-- {
        if words[i] != words[0] {
            a = i + 1
            break
        }
    }
    for i := 1; i < n - 1; i++ {
        if words[i] != words[n - 1] {
            b = n - i
            break
        }
    }

    return max(a, b)
}