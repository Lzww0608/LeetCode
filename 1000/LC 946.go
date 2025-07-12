func validateStackSequences(pushed []int, popped []int) bool {
    st := []int{}
    n := len(popped)
    j := 0
    for _, x := range pushed {
        st = append(st, x)
        for j < n && len(st) > 0 && st[len(st)-1] == popped[j] {
            j++
            st = st[:len(st)-1]
        }
    }

    return j == n 
}