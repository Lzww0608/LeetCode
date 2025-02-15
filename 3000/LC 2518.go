const MOD int = 1_000_000_007
func countPartitions(nums []int, k int) int {
    sum := 0
    for _, x := range nums {
        sum += x
    }
    if sum < k * 2 {
        return 0
    }

    f := make([]int, k)
    ans := 1
    f[0] = 1
    for _, x := range nums {
        ans = (ans * 2) % MOD
        for i := k - 1; i >= x; i-- {
            f[i] = (f[i] + f[i-x]) % MOD
        }
    }

    for _, x := range f {
        ans = (ans - x * 2 % MOD + MOD) % MOD
    }

    return ans
}
