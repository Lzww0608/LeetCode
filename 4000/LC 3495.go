
import "math/bits"
func minOperations(queries [][]int) int64 {
    ans := 0
    for _, q := range queries {
        ans += (calc(q[1]) - calc(q[0] - 1) + 1) / 2
    }

    return int64(ans)
}

func calc(n int) (res int) {
    m := bits.Len(uint(n))
    for i := 1; i < m; i++ {
        res += ((i + 1) / 2) << (i - 1) 
    }

    return res + (m + 1) / 2 * (n + 1 - (1 << m >> 1))
}