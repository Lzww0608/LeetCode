const MOD int = 1_000_000_007
func goodBinaryStrings(minLength int, maxLength int, oneGroup int, zeroGroup int) (ans int) {
    f := make([]int, maxLength + 1)
    f[0] = 1
    for i := 0; i < maxLength; i++ {
        if f[i] != 0 {
            if j := i + oneGroup; j <= maxLength {
                f[j] += f[i]
                f[j] %= MOD
            }
            if j := i + zeroGroup; j <= maxLength {
                f[j] += f[i]
                f[j] %= MOD
            }
        }
    }

    for _, x := range f[minLength:] {
        ans = (ans + x) % MOD
    }

    return
}