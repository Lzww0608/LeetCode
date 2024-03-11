func nextGreaterElements(a []int) []int {
    n := len(a)
    ans := make([]int, n)
    for i := range ans {
        ans[i] = -1
    }

    st := []int{}
    for i := 0; i < 2 * n - 1; i++ {
        for len(st) > 0 && a[i % n] > a[st[len(st)-1]] {
            ans[st[len(st)-1]] = a[i % n]
            st = st[:len(st)-1]
        }
        
        st = append(st, i % n)
    }
    return ans
}