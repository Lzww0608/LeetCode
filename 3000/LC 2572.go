const (
    MOD int = 1_000_000_007 
    N int = 30
)

func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = res * a % MOD
        }
        a = a * a % MOD
        r >>= 1
    }

    return res
}

func squareFreeSubsets(nums []int) (ans int) {
    primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

    freq := make([]int, N + 1)
    for _, x := range nums {
        freq[x]++
    }

    ans = quickPow(2, freq[1]) - 1
    f := make([]int, 1 << len(primes))
    f[0] = ans + 1
    
    for i := 2; i <= N; i++ {
        if freq[i] == 0 {
            continue
        }

        sf := true 
        mask := 0
        for j, x := range primes {
            if i % (x * x) == 0 {
                sf = false
                break
            }

            if i % x == 0 {
                mask |= 1 << j 
            }
        }

        if sf {
            for j := 1; j < len(f); j++ {
                if j & mask == mask {
                    f[j] = (f[j] + f[j ^ mask] * freq[i]) % MOD
                }
            }
        }
    }

    for _, x := range f[1:] {
        ans = (ans + x) % MOD
    }

    return
}