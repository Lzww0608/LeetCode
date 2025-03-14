
import "math"
const N int = 1_000_001

// var primes []int
// var isPrimes []bool

// func init() {
//     isPrimes = make([]bool, N)
//     for i := 2; i < N; i++ {
//         if !isPrimes[i] {
//             primes = append(primes, i)
//             for j := i * i; j < N; j += i {
//                 isPrimes[j] = true
//             }
//         }
//     }
// }

func splitArray(nums []int) int {
    pre := 0
    f := make([]int, N)
    for i := range f {
        f[i] = math.MaxInt32
    }

    for _, x := range nums {
        now := math.MaxInt32
        for y := 2; y * y <= x; y++ {
            if x % y == 0 {
                f[y] = min(f[y], 1 + pre)
                now = min(now, f[y])
                for x % y == 0 {
                    x /= y
                }
            } 
        }
        if x > 1 {
            f[x] = min(f[x], 1 + pre)
            now = min(now, f[x])
        }
        pre = now
    }

    return pre
}