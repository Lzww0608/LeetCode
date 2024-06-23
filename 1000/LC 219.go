func containsNearbyDuplicate(nums []int, k int) bool {
    m := map[int]int{}
    for i, x := range nums {
        if j, ok := m[x]; ok {
            if i - j <= k {
                return true
            }
        }
        m[x] = i
    }

    return false
}

