
import (
	"math"
	"math/bits"
)
func minimumXORSum(a []int, b []int) (ans int) {
    n := len(a)
    f := make([]int, 1 << n)
    for i := range f {
        f[i] = math.MaxInt32
    }
    f[0] = 0

    for mask := 0; mask < (1 << n); mask++ {
        i := bits.OnesCount(uint(mask))
        for j := 0; j < n; j++ {
            if mask & (1 << j) == 0 {
                f[mask|(1<<j)] = min(f[mask|(1<<j)], f[mask] + (a[i] ^ b[j]))
            }
        }
    }

    return f[(1 << n) - 1]
}