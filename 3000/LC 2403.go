
import "math"
func minimumTime(power []int) int64 {
    n := len(power)
    m := 1 << n
    f := make([]int, m)
    for i := range f {
        f[i] = math.MaxInt64 / 2
    }
    f[0] = 0
    
    for s := 1; s < m; s++ {
        for i := 0; (1 << i) <= s; i++ {
            if (s >> i) & 1 == 1 {
                t := bits.OnesCount(uint(s))
                x := (power[i] + t - 1) / t
                f[s] = min(f[s], f[s^(1<<i)] + x)
            }
        }
    }

    return int64(f[m-1])
}