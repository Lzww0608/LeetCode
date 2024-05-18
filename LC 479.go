func largestPalindrome(n int) int {
    const MOD int = 1337

    if n == 1 {
        return 9
    }

    mx := 1
    for i := 0; i < n; i++ {
        mx *= 10
    }
    mx--

    for i := mx; i > 0; i-- {
        x := i
        for d := i; d > 0; d /= 10 {
            x = x * 10 + d % 10
        }
        for j := mx; j * j >= x; j-- {
            if x % j == 0 {
                return x % MOD
            }
        }
    }

    return -1
}