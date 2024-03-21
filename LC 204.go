func countPrimes(n int) (ans int) {
    f := make([]int, n)
    for i := 2; i < n; i++ {
        if f[i] == 1 {
            continue
        }
        f[i] = 1
        ans++
        for j := i * i; j < n; j += i {
            f[j] = 1
        }
    }

    return 
}