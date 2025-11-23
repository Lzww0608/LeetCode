func minBitwiseArray(nums []int) []int {
    for i, x := range nums {
        if x & 1 == 0 {
            nums[i] = -1
            continue
        } 
        d := 0
        for x & (1 << d) != 0 {
            d++
        }
        mask := (1 << d) - 1
        nums[i] = (x &^ mask) | (mask >> 1)
    }

    return nums
}