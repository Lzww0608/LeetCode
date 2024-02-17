var MOD int = int(1e9 + 7)
func sumSubseqWidths(nums []int) int {
    n := len(nums)
    sort.Ints(nums)
    ans, pow2 := 0, 1
    for i := 0; i < n; i++ {
        ans += (nums[i] - nums[n-1-i]) * pow2
        pow2 = pow2 * 2 % MOD
        ans = (ans + MOD) % MOD
    }
    return ans
}