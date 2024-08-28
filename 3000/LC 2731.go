const MOD int = 1_000_000_007
func sumDistance(nums []int, s string, d int) (ans int) {
    for i := range nums {
        if s[i] == 'R' {
            nums[i] = (nums[i] + d)
        } else {
            nums[i] = (nums[i] - d) 
        }
    }

    sort.Ints(nums)
    pre := 0
    for i, x := range nums {
        ans = (ans + i * x - pre) % MOD
        pre += x
    }

    return 
}