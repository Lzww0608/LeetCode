const MOD int = 1_000_000_007
func getSum(nums []int) (ans int) {
    inc, a := nums[0], 1
    des, b := nums[0], 1 
    ans += nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] + 1 {
            des, b = nums[i], 1 
            inc, a = inc + (a + 1) * nums[i], a + 1
            ans = (ans + inc) % MOD
        } else if nums[i] == nums[i-1] - 1 {
            des, b = des + (b + 1) * nums[i], b + 1
            inc, a = nums[i], 1 
            ans = (ans + des) % MOD
        } else {
            inc, a = nums[i], 1 
            des, b = nums[i], 1
            ans = (ans + nums[i]) % MOD
        }
    } 

    return
}