func arraySign(nums []int) int {
    f := 1
    for _, x := range nums {
        if x < 0 {
            f = -f
        } else if x == 0 {
            return 0
        }
    }

    return f
}