func hasTrailingZeros(nums []int) bool {
    cnt := 0
    for _, x := range nums {
        if x & 1 == 0 {
            cnt++
        }
    }

    return cnt >= 2
}