const MOD int = 1_000_000_007
func sumSubarrayMins(arr []int) (ans int) {
    st := []int{-1}
    arr = append(arr, -1)
    for i, x := range arr {
        for len(st) > 1 && x < arr[st[len(st)-1]] {
            mid := st[len(st)-1]
            st = st[:len(st)-1]
            j := st[len(st)-1]
            y := arr[mid]
            ans = (ans + (i - mid) * (mid - j) * y % MOD) % MOD
        }
        st = append(st, i)
    }

    return 
}