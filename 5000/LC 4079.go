const N int = 1_000_001

var a []int 
var f []int
var isPrime [N]bool

func init() {
    for i := 2; i < N; i++ {
        if isPrime[i] {
            continue
        }
        a = append(a, i)
        if len(f) > 0 {
            f = append(f, f[len(f) - 1] + i)
        } else {
            f = append(f, i)
        }
        for j := i * i; j < N; j += i {
            isPrime[j] = true
        }
    }
}

func maxPrimes(n int, s int) []int {
    i := sort.SearchInts(a, n + 1)
    j := sort.SearchInts(f, s + 1)

    return a[:min(i, j)]
}