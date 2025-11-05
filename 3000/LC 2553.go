func separateDigits(nums []int) (ans []int) {
    a := []int{}
    for _, x := range nums {
        a = a[:0]
        for x > 0 {
            a = append(a, x % 10)
            x /= 10
        }
        slices.Reverse(a)
        ans = append(ans, a...)
    }

    return
}