func missingNumber(nums []int) int {
    x, i := 0, 0
    for _, y := range nums {
        x ^= i ^ y
        i++
    }
    return x ^ i
}
