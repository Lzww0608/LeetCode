func maximumStrongPairXor(nums []int) (ans int) {
    sort.Ints(nums)
    n := len(nums)
    d := bits.Len(uint(nums[n-1])) - 1
    m := make(map[int]int)
    mask := 0

    for i := d; i >= 0; i-- {
        clear(m)
        mask |= 1 << i 
        cur := ans | (1 << i)
        for _, x := range nums {
            t := x & mask
            if y, ok := m[t ^ cur]; ok && y * 2 >= x {
                ans = cur
                break
            }
            m[t] = x
        }
    }
    return
}