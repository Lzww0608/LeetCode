const MOD int = 1_000_000_007
func beautifulPartitions(s string, k int, minLength int) int {
    n := len(s)
    if k * minLength > n || !isPrime(s[0]) || isPrime(s[n-1]) {
        return 0
    }
    f := make([][]int, k + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
    }
    f[0][0] = 1

    check := func(j int) bool {
        return j == 0 || j == n || !isPrime(s[j-1]) && isPrime(s[j])
    }
    
    for i := 1; i <= k; i++ {
        sum := 0
        for j := i * minLength; j + (k - i) * minLength <= n; j++ {
            if check(j - minLength) {
                sum = (sum + f[i-1][j-minLength]) % MOD
            }

            if check(j) {
                f[i][j] = sum
            }
        }
    }

    return f[k][n]
}


func isPrime(c byte) bool {
    if c == '2' || c == '3' || c == '5' || c == '7' {
        return true
    }

    return false
}