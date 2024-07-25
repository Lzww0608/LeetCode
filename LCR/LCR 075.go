func relativeSortArray(arr1 []int, arr2 []int) []int {
    m := map[int]int{}
    for i, x := range arr2 {
        m[x] = i
    }

    sort.Slice(arr1, func(i, j int) bool {
        _, ok1 := m[arr1[i]]
        _, ok2 := m[arr1[j]]
        if ok1 {
            if ok2 {
                return m[arr1[i]] < m[arr1[j]]
            } else {
                return true
            }
            
        } else {
            if !ok2 {
                return arr1[i] < arr1[j]
            } else {
                return false
            }
        }
    })

    return arr1
}