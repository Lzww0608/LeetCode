func findPrimePairs(n int) (ans [][]int) {
    isPrimes := make([]int, n + 1)
    primes := make(map[int]bool)
    for i := 2; i <= n; i++ {
        if isPrimes[i] == 1 {
            continue
        }
        primes[i] = true
        if primes[n - i] {
            ans = append(ans, []int{n - i, i})
        }

        for j := i * i; j <= n; j += i {
            isPrimes[j] = 1
        }
    }

    m := len(ans)
    for i, j := 0, m - 1; i < j; i, j = i + 1, j - 1 {
        ans[i], ans[j] = ans[j], ans[i]
    }

    return
}