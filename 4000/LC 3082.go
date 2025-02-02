const MOD int = 1_000_000_007
func sumOfPower(nums []int, k int) int {
    n := len(nums)
    f := make([][]int, k + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
    }

    f[0][0] = 1
    for _, x := range nums {
        for j := k; j >= x; j-- {
            for t := n - 1; t >= 0; t-- {
                f[j][t+1] = (f[j][t+1] + f[j-x][t]) % MOD
            }
        }
    }

    ans := 0
    for i := 1; i <= n; i++ {
        ans = (ans + f[k][i] * quickPow(2, n - i)) % MOD
    }
    return ans
}

func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = (res * a) % MOD
        }
        a = (a * a) % MOD
        r >>= 1
    }
    return res
}