const (
    MOD int = 1_000_000_007
    N int = 30
)
func numberOfGoodSubsets(nums []int) (ans int) {
    primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
    
    freq := make([]int, N + 1)
    for _, x := range nums {
        freq[x]++
    }

    f := make([]int, 1 << len(primes))
    f[0] = 1
    for i := 0; i < freq[1]; i++ {
        f[0] = f[0] * 2 % MOD
    }

    for i := 2; i <= N; i++ {
        if freq[i] == 0 {
            continue
        }

        check := true
        mask := 0
        for j, x := range primes {
            if i % (x * x) == 0 {
                check = false
                break
            } 

            if i % x == 0 {
                mask |= 1 << j
            }
        }
        if check {
            for j := (1 << len(primes)) - 1; j > 0; j-- {
                if j & mask == mask {
                    f[j] = (f[j] + f[j ^ mask] * freq[i]) % MOD
                }
            }
        }
    }

    for _, x := range f[1:] {
        ans += x
    }

    return ans % MOD
}