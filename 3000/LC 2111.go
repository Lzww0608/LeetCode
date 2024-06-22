func kIncreasing(arr []int, k int) (ans int) {
    n := len(arr)
    st := make([][]int, min(n, k))
    for i := range st {
        st[i] = append(st[i], math.MinInt32)
    }
    for i, x := range arr {
        if x < st[i%k][len(st[i%k])-1] {
            pos := sort.SearchInts(st[i%k], x + 1)
            st[i%k][pos] = x
        } else {
            st[i%k] = append(st[i%k], x)
        }
    }
    ans = n
    for _, v := range st {
        ans -= len(v) - 1
    }

    return 
}
