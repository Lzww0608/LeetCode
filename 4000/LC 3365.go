func isPossibleToRearrange(s string, t string, k int) bool {
    m := make(map[string]int)
    d := len(s) / k
    for i := 0; i < k; i++ {
        m[s[i * d : d * (i + 1)]]++
        m[t[i * d : d * (i + 1)]]--
    }

    for _, v := range m {
        if v != 0 {
            return false
        }
    }

    return true
}