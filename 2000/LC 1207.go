func uniqueOccurrences(arr []int) bool {
    cnt := make(map[int]int)
    for _, x := range arr {
        cnt[x]++
    }

    m := make(map[int]struct{})
    for _, x := range cnt {
        if _, ok := m[x]; ok {
            return false
        } 
        m[x] = struct{}{}
    }

    return true
}