const N int = 1_000_005

var isPrime [N]int

var primes []int

func init() {
    for i := 2; i < N; i++ {
        if isPrime[i] == 1 {
            continue
        }
        primes = append(primes, i)
        for j := i * i; j < N; j += i {
            isPrime[j] = 1
        }
    }
}

func closestPrimes(left int, right int) []int {
    ans := []int{-1, -1}

    for i := 1; i < len(primes); i++ {
        if primes[i-1] >= left && primes[i] <= right && (ans[0] == -1 || primes[i] - primes[i-1] < ans[1] - ans[0]) {
            ans = []int{primes[i-1], primes[i]}
            if ans[1] - ans[0] == 1 {
                return ans
            }
        }
    }

    return ans
}