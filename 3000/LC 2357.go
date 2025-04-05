func minimumOperations(nums []int) int {
    m := make(map[int]bool)
    for _, x := range nums {
        if x != 0 {
            m[x] = true
        }
    }

    return len(m)
}