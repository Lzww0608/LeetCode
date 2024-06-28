func validateBookSequences(putIn []int, takeOut []int) bool {
    n, i, j := len(putIn), 0, 0
    st := []int{-1}
    for i < n && j < n {
        for i < n && len(st) > 0 && st[len(st)-1] != takeOut[j] {
            st = append(st, putIn[i])
            i++
        } 
        for j < n && len(st) > 0 && takeOut[j] == st[len(st)-1] {
            st = st[:len(st)-1]
            j++
        }
    }
    
    for j < n && len(st) > 0 && st[len(st)-1] == takeOut[j] {
        j++
        st = st[:len(st)-1]
    }

    return i == n && j == n
}
