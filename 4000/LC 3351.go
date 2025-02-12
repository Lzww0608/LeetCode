const MOD int = 1_000_000_007
func sumOfGoodSubsequences(nums []int) (ans int) {
    f := make(map[int]int)
    g := make(map[int]int)
    for _, x := range nums {
        f[x] = (f[x] + x * (1 + g[x-1] + g[x+1]) % MOD + f[x-1] + f[x+1]) % MOD
        g[x] = (1 + g[x] + g[x-1] + g[x+1]) % MOD
    }

    for _, x := range f {
        ans = (ans + x) % MOD
    }

    return
}