
import "math"
func nthSuperUglyNumber(n int, primes []int) int {
    m := len(primes)
    ugly := make([]int, n)
    p := make([]int, m)
    ugly[0] = 1

    for i := 1; i < n; i++ {
        mn := math.MaxInt
        for j, x := range primes {
            mn = min(mn, x * ugly[p[j]])
        }

        ugly[i] = mn
        for j, x := range primes {
            if mn == x * ugly[p[j]] {
                p[j]++
            }
        }
    }

    return ugly[n - 1]
}