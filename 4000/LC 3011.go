func canSortArray(nums []int) bool {
    n := len(nums)
    mx := 0
    r := 1
    cur := bits.OnesCount(uint(nums[0]))
    nowMax, nowMin := nums[0], nums[0]
    for r < n {
        x := bits.OnesCount(uint(nums[r]))
        if x == cur {
            nowMax = max(nowMax, nums[r])
            nowMin = min(nowMin, nums[r])
        } else {
            mx = nowMax
            nowMax, nowMin = nums[r], nums[r]
            cur = x
        }
        if nowMin < mx {
            return false
        }  
        r++
    }

    return true
}