func sumAndMultiply(n int) int64 {
    x, sum := 0, 0

    for p := 1; n > 0; n /= 10 {
        if n % 10 != 0 {
            x += p * (n % 10)
            p *= 10
            sum += n % 10
        }
    }

    return int64(x * sum)
}