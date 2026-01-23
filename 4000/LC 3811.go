const MOD int = 1_000_000_007
func alternatingXOR(nums []int, target1 int, target2 int) int {
    n := len(nums)
    f := make(map[int]int)
    g := make(map[int]int)
    g[0] = 1
    xor := 0
    for i, x := range nums {
        xor ^= x 

        a, b := g[xor ^ target1], f[xor ^ target2]
        if i + 1 == n {
            return (a + b) % MOD
        }
        f[xor] = (f[xor] + a) % MOD
        g[xor] = (g[xor] + b) % MOD
    }

    return -1
}