func evenNumberBitwiseORs(nums []int) (ans int) {
    for _, x := range nums {
        if x & 1 == 0 {
            ans |= x
        }
    }

    return
}