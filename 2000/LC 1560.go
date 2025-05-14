func mostVisited(n int, rounds []int) (ans []int) {
    m := len(rounds)
    start, end := rounds[0], rounds[m-1]
    if start <= end {
        for i := start; i <= end; i++ {
            ans = append(ans, i)
        }
    } else {
        for i := 1; i <= end; i++ {
            ans = append(ans, i)
        }
        for i := start; i <= n; i++ {
            ans = append(ans, i)
        }
    }

    return
}