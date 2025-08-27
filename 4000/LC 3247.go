const MOD int = 1_000_000_007
func subsequenceCount(nums []int) (ans int) {
    odd, even := 0, 0
    for _, x := range nums {
        if x & 1 == 1 {
            odd++
        } else {
            even++
        }
    }
    if odd == 0 {
        return 0
    }

    return quickPow(2, odd - 1) * quickPow(2, even) % MOD
}

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