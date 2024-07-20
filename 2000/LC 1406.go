func stoneGameIII(stoneValue []int) string {
    n := len(stoneValue)
    f := make([]int, n)
    for i := range f {
        f[i] = math.MinInt
    }

    for i := n - 1;i >= 0; i-- {
        s := 0
        for j := i + 1; j <= i + 3; j++ {
            if j <= n {
                s += stoneValue[j-1]
                if j == n {
                    f[i] = max(f[i], s)
                } else {
                    f[i] = max(f[i], s - f[j])
                }
            }
        }
    }

    if f[0] > 0 {
        return "Alice"
    } else if f[0] < 0 {
        return "Bob"
    }

    return "Tie"
}