var m = map[int]int{}
func getKth(l int, h int, k int) int {

    a := make([]int, h - l + 1)
    m[1] = 0
    for i := range a {
        a[i] = l + i
    }
    sort.Slice(a, func(i, j int) bool {
        if getValue(a[i]) != getValue(a[j]) {
            return getValue(a[i]) < getValue(a[j])
        }
        return a[i] < a[j]
    })
    return a[k-1]
}



func getValue(x int) int {
    if v, ok := m[x]; ok {
        return v
    }
    if x & 1 == 1 {
        m[x] = getValue(3 * x + 1) + 1
    } else {
        m[x] = getValue(x / 2) + 1
    }
    return m[x]
} 
