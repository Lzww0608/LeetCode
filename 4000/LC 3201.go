func maximumLength(nums []int) int {
    a, b, c, d := 0, 0, 0, 0
    for _, x := range nums {
        if x & 1 == 0 {
            b++
            d = c + 1
        } else {
            a++
            c = d + 1
        }
    }

    return max(a, b, c, d)
}