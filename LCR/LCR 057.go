func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
    m := map[int]int{}

    getID := func(a, b int) int {
        if a >= 0 {
            return a / b
        }
        
        return (a + 1) / b - 1
    }

    for i, x := range nums {
        id := getID(x, t + 1)
        
        if _, ok := m[id]; ok {
            return true
        }
        if v, ok := m[id-1]; ok && abs(v - x) <= t {
            return true
        }
        if v, ok := m[id+1]; ok && abs(v - x) <= t {
            return true
        }

        m[id] = x
        if i >= k {
            delete(m, getID(nums[i-k], t + 1))
        }
    }

    return false
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}