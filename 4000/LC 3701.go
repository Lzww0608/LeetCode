func alternatingSum(nums []int) int {
    sum := 0
    for i, x := range nums {
        if i & 1 == 0 {
            sum += x
        } else {
            sum -= x
        }
    }

    return sum
}