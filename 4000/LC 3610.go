const N int = 1_000

var primes []int

func init() {
    isPrime := make(map[int]bool)
    for i := 2; i <= N; i++ {
        if isPrime[i] {
            continue
        }
        primes = append(primes, i)
        for j := i * i; j <= N; j += i {
            isPrime[j] = true 
        }
    }
}

func minNumberOfPrimes(n int, m int) int {
    f := make([]int, n + 1)
    for i := range f {
        f[i] = n
    }
    f[0] = 0
    for i := 2; i <= n; i++ {
        for _, x := range primes[:min(len(primes), m)] {
            if x > i {
                break
            }
            f[i] = min(f[i], f[i - x] + 1)
        }
    }

    if f[n] >= n {
        return -1
    }
    return f[n]
}