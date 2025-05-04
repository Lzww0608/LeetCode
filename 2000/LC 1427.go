func stringShift(s string, shift [][]int) string {
    id := 0
    n := len(s)
    for _, v := range shift {
        if v[0] == 0 {
            id += v[1]
        } else {
            id += n * v[1] - v[1] 
        }
        id %= n
    }

    return s[id:] + s[:id]
}