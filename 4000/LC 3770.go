const N int = 500_001

var isPrime [N]bool 

func init() {
    for i := 2; i < N; i++ {
        if isPrime[i] {
            continue
        }

        for j := i * i; j < N; j += i {
            isPrime[j] = true
        }
    }
}

func largestPrime(n int) (ans int) {
    cur := 0
    for i := 2; i <= n; i++ {
        if isPrime[i] {
            continue
        }

        cur += i 
        if cur > n {
            break
        }
        if !isPrime[cur] {
            ans = max(ans, cur)
        }
    }

    return
}