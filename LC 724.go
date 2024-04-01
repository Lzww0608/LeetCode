func pivotIndex(nums []int) int {
    suf := 0
    for _, x := range nums {
        suf += x
    }

    pre := 0
    for i := range nums {
        suf -= nums[i]
        if pre == suf {
            return i
        }
        pre += nums[i]
    }

    return -1
}