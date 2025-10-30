func pivotInteger(n int) int {
    // x * (x + 1) / 2
    // n * (n + 1) / 2 - x * (x - 1) / 2
    // n * (n + 1) / 2 == x * x

    x := int(math.Sqrt(float64(n * (n + 1) / 2)))
    if x * x == n * (n + 1) / 2 {
        return x
    }

    return -1
}