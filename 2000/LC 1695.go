func maximumUniqueSubarray(nums []int) (ans int) {
    n := len(nums)
    sum := 0
    l, r := 0, 0
    m := map[int]int{}
    for r < n {
        x := nums[r]
        if m[x] > 0 {
            for l < r && m[x] > 0 {
                m[nums[l]]--
                sum -= nums[l]
                l++
            }
        } 
        m[x]++
        sum += x
        ans = max(ans, sum)
        r++
    }
    return
}
