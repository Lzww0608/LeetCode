const MOD int = 1_000_000_007
func countSpecialSubsequences(nums []int) int {
    n := len(nums)
    f := make([][3]int, n + 1)
    for i, x := range nums {
        f[i+1] = f[i]
        if x == 0 {
            f[i+1][0] = (2 * f[i][0] + 1) % MOD
        } else if x == 1 {
            f[i+1][1] = (2 * f[i][1] + f[i][0]) % MOD
        } else {
            f[i+1][2] = (2 * f[i][2] + f[i][1]) % MOD
        }
    }

    return f[n][2]
}