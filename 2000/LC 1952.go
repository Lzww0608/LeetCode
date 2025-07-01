func isThree(n int) bool {
    isPrime := make([]bool, n)
    for i := 2; i * i <= n; i++ {
        if isPrime[i] {
            continue
        }
        if i * i == n {
            return true
        }

        for j := i * i; j * j <= n; j += i {
            isPrime[j] = true
        }
    }

    return false
}