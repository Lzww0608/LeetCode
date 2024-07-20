func findMaximumXOR(nums []int) (ans int) {
    highBit := bits.Len(uint(slices.Max(nums))) - 1
    exist := map[int]bool{}
    mask := 0
    for i := highBit; i >= 0; i-- {
        clear(exist)
        mask |= 1 << i
        cur := ans | (1 << i)
        for _, x := range nums {
            x &= mask
            if exist[cur ^ x] {
                ans = cur
                break
            }
            exist[x] = true
        }
    }

    return
}