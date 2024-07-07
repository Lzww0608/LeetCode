func findMiddleIndex(nums []int) int {
    sum := 0
    for _, x := range nums {
        sum += x
    }

    cur := 0
    for i, x := range nums {
        if cur == sum - x {
            return i
        }
        cur += x
        sum -= x
    }

    return -1
}