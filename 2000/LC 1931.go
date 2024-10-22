const MOD int = 1_000_000_007

func colorTheGrid(m int, n int) (ans int) {
    m <<= 1
    valid := []int{}
    for i := 0; i < (1 << m); i++ {
        f := true
        pre := 3
        for j := 0; j < m; j += 2 {
            if t := (i >> j) & 3; t != 3 && t != pre {
                pre = t
            } else {
                f = false
                break
            }
        }
        if f {
            valid = append(valid, i)
        }
    }

    sz := len(valid)
    to := make([][]int, sz)
    for i := range valid {
        for j := i + 1; j < sz; j++ {
            f := true
            for k := 0; k < m; k += 2 {
                if a, b := (valid[i] >> k) & 3, (valid[j] >> k) & 3; a == b {
                    f = false
                    break
                }
            }
            if f {
                to[i] = append(to[i], j)
                to[j] = append(to[j], i)
            }
        }
    }

    f := make([]int, sz)
    for i := range f {
        f[i] = 1
    }
    for i := 1; i < n; i++ {
        tmp := f 
        f = make([]int, sz)
        for j, k := range tmp {
            for _, t := range to[j] {
                f[t] = (f[t] + k) % MOD
            }
        }
    }

    for _, x := range f {
        ans = (ans + x) % MOD
    }

    return 
}