func findKOr(nums []int, k int) (ans int) {
    for i := 0; i <= 31; i++ {
        mask, cnt := 1 << i, 0
        for _, x := range nums {
            if mask & x == mask {
                cnt++
            }
        } 
        if cnt >= k {
            ans |= mask
        }
    }
    return 
}