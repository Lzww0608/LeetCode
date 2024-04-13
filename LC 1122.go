func relativeSortArray(arr1 []int, arr2 []int) []int {
    m := map[int]int{}
    for i, x := range arr2 {
        m[x] = i
    }

    sort.Slice(arr1, func(i, j int) bool {
        x, ok1 := m[arr1[i]]
        y, ok2 := m[arr1[j]]
        if ok1 && ok2  {
            return x < y
        } else if !ok1 && !ok2 {
            return arr1[i] < arr1[j]
        } else if ok1 {
            return true
        } else {
            return false
        }
    })

    return arr1
}