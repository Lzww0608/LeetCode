
import "math/bits"
func minNumberOfSemesters(n int, relations [][]int, k int) int {
    f := make([]int, 1 << n)
    pre := make([]int, n)
    for _, v := range relations {
        a, b := v[0] - 1, v[1] - 1
        pre[b] |= 1 << a
    }

    for i := range f {
        f[i] = n
    }
    f[0] = 0
    for s := 0; s < (1 << n); s++ {
        next := 0
        for i := 0; i < n; i++ {
            if (s & pre[i]) == pre[i] {
                next |= 1 << i
            }
        }
        next &= ^s
        if bits.OnesCount(uint(next)) <= k {
            f[s|next] = min(f[s|next], f[s] + 1)
            continue
        }
        for i := next; i > 0; i = (i - 1) & next {
            if bits.OnesCount(uint(i)) == k {
                f[s|i] = min(f[s|i], f[s] + 1)
            }
        }

    }

    return f[1<<n-1]
}