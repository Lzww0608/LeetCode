func minSteps(n int) int {
    if n == 1 {
        return 0
    }
    i := 2
    for n % i != 0 {
        i++
    }

    return i + minSteps(n / i)
}