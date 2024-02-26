func singleNumber(nums []int) int {
    var ans int32 = 0
    for i := 0; i < 32; i++ {
        var t int32 = 1 << i
        k := 0
        for _, x := range nums {
            if t & int32(x) != 0 {
                k++
            }
        }
        if k % 3 != 0 {
            ans |= t
        }
    }
    return int(ans)
}