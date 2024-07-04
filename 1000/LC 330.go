func minPatches(nums []int, n int) (ans int) {
    s, i := 1, 0
    m := len(nums)
    for s <= n {
        if i < m && nums[i] <= s {
            s += nums[i]
            i++
        } else {
            s <<= 1
            ans++
        }
    }

    return
}