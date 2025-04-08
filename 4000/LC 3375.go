func minOperations(nums []int, k int) int {
    m := make(map[int]bool)
    for _, x := range nums {
        if x < k {
            return -1
        } 
        m[x] = true
    }

    delete(m, k)

    return len(m)
}