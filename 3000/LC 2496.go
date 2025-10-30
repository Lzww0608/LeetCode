func maximumValue(strs []string) (ans int) {
    for _, s := range strs {
        x, err := strconv.Atoi(s)
        if err == nil {
            ans = max(ans, x)
        } else {
            ans = max(ans, len(s))
        }
    }

    return
}