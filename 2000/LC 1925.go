func countTriples(n int) (ans int) {
    check := func(x int) bool {
        y := int(math.Sqrt(float64(x)))
        return y <= n && y * y == x
    }

    for i := 1; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if check(i * i + j * j) {
                ans += 2
            }
        }
    }

    return
}


