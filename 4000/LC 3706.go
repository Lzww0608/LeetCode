func maxDistance(s []string) int {
    a, b := 0, 0
    n := len(s)
    for i := n - 1; i > 0; i-- {
        if s[i] != s[0] {
            a = i + 1 
            break 
        }
    }

    for i := 1; i < n - 1; i++ {
        if s[i] != s[n - 1] {
            b = n - i 
            break
        }
    }

    return max(a, b)
}