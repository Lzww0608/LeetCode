func goodIndices(s string) (ans []int) {
    for i := range s {
        t := strconv.Itoa(i)
        n := len(t)
        if i >= n - 1 && s[i - n + 1:i + 1] == t {
            ans = append(ans, i)
        }
    }

    return
}