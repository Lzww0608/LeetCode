var MOD int = int(1e9 + 7)
func sumOfPower(nums []int) int {
    sort.Ints(nums)
    pre, ans := 0, 0
    for _, x := range nums {
        ans = (ans + x * x % MOD * (x + pre) % MOD) % MOD
        pre = (pre * 2 + x) % MOD
    }
    return ans
}
