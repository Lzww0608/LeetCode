func pivotIndex(nums []int) int {
    sum := 0
    for _, x := range nums {
        sum += x
    }

    cur := 0
    for i, x := range nums {
        sum -= x
        if cur == sum {
            return i
        }
        cur += x
    }

    return -1
}