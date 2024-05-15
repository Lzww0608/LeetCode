func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
    m := make(map[int]int)
    
    getID := func(a, b int) int {
        if a >= 0 {
            return a / b
        } 
        return (a + 1) / b - 1
    }

    for i, x := range nums {
        id := getID(x, valueDiff + 1)
        if _, ok := m[id]; ok {
            return true
        } 
        if v, ok := m[id-1]; ok && abs(v - x) <= valueDiff {
            return true
        } 
        if v, ok := m[id+1]; ok && abs(v - x) <= valueDiff {
            return true
        }

        m[id] = x
        if i >= indexDiff {
            delete(m, getID(nums[i-indexDiff], valueDiff + 1))
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