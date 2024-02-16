func superEggDrop(k int, n int) int {
    t := 1;
    for dfs(k, t) < n + 1 {
        t++
    }
    return t
}

func dfs(k, t int) int {
    if k == 1 || t == 1 {
        return t + 1
    }
    return dfs(k-1, t-1) + dfs(k, t-1)
}