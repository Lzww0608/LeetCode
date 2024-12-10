const N int = 1001
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

func primeSubOperation(nums []int) bool {
    pre := 0
    for _, x := range nums {
        pos := sort.SearchInts(primes, x - pre) - 1
        if pos >= 0 {
            x -= primes[pos]
        }
        if x <= pre {
            return false
        }
        pre = x
    }
    return true
}