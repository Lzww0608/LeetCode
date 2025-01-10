const MOD int = 1_000_000_007
func purchasePlans(nums []int, target int) (ans int) {
    n := len(nums);
    sort.Ints(nums)
    l, r := 0, n - 1
    for l < r {
        if nums[l] + nums[r] > target {
            r--
        } else {
            ans = (ans + r - l) % MOD
            l++
        }
    }

    return 
}