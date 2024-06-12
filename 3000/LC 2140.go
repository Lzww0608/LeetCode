func mostPoints(questions [][]int) int64 {
    n := len(questions)
    f := make([]int64, n + 1)
    for i := n - 1; i >= 0; i-- {
        a, b := questions[i][0], questions[i][1]
        if i + b + 1 < n {
            f[i] = int64(a) + f[i + b + 1]
        } else {
            f[i] = int64(a)
        }
        f[i] = max(f[i+1], f[i])
    }

    return f[0]
}