func subarraysWithKDistinct(nums []int, k int) (ans int) {
    m := map[int]int{}
    mp := map[int]int{}
    n := len(nums)
    l, r := 0, 0
    pre := 0
    for r < n {
        m[nums[r]]++
        mp[nums[r]]++
        for len(mp) >= k {
            mp[nums[pre]]--
            if mp[nums[pre]] == 0 {
                delete(mp, nums[pre])
            }
            pre++
        }

        for l < r && len(m) > k {
            m[nums[l]]--
            if m[nums[l]] == 0 {
                delete(m, nums[l])
            }
            l++
        }
            
        if len(m) == k {
            ans += pre - l
        }
        r++
    }

    return
}