const MOD int = 1_000_000_007
func maxFrequencyScore(nums []int, k int) (ans int) {
    st := make(map[int][]int)
    n := len(nums)
    sum := 0
    for l, r := 0, 0; r < n; r++ {
        if _, ok := st[nums[r]]; !ok || len(st[nums[r]]) == 0 {
            st[nums[r]] = []int{nums[r]}
            sum = (sum + nums[r]) % MOD
        } else {
            if len(st[nums[r]]) > 0 {
                sum = (sum - st[nums[r]][len(st[nums[r]])-1] + MOD) % MOD
            }
            st[nums[r]] = append(st[nums[r]], nums[r] * st[nums[r]][len(st[nums[r]])-1] % MOD)
            sum = (sum + st[nums[r]][len(st[nums[r]])-1]) % MOD
        }
        
        if r - l + 1 == k {
            ans = max(ans, sum)
            t := st[nums[l]][len(st[nums[l]])-1]
            st[nums[l]] = st[nums[l]][:len(st[nums[l]])-1]
            sum = (sum - t + MOD) % MOD
            if len(st[nums[l]]) > 0 {
                sum = (sum + st[nums[l]][len(st[nums[l]])-1]) % MOD
            }
            l++
        }
    }

    return 
}

