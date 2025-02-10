const n int = 12
func minTransfers(a [][]int) int {
    m := 1 << n 
    f := make([]int, m)
    cnt := make([]int, n)
    for _, v := range a {
        cnt[v[0]] -= v[2]
        cnt[v[1]] += v[2]
    }

    for mask := 1; mask < m; mask++ {
        sum := 0
        for i := 0; i < n; i++ {
            if (mask >> i) & 1 == 1 {
                sum += cnt[i]
            }
        }
        
        if sum != 0 {
            f[mask] = math.MaxInt32 / 2
        } else {
            f[mask] = bits.OnesCount(uint(mask)) - 1
            for j := mask; j > 0; j = (j - 1) & mask {
                f[mask] = min(f[mask], f[j] + f[mask ^ j])
            }
        }
    }

    return f[m-1]
}