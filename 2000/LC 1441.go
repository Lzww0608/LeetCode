func buildArray(target []int, n int) (ans []string) {
    m := len(target)
    for i, j := 1, 0; i <= n && j < m; i++ {
        ans = append(ans, "Push")
        if i != target[j] {
            ans = append(ans, "Pop")
        } else {
            j++
        }
    }

    return 
}