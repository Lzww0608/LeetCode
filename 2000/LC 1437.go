func kLengthApart(nums []int, k int) bool {
    cnt := k
    for _, x := range nums {
        if x == 1 {
            if cnt < k {
                return false
            }
            cnt = 0
        } else {
            cnt++
        }
    }

    return true
}