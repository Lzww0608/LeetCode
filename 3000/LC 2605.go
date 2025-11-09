const N int = 10
func minNumber(a []int, b []int) int {
    cnt1 := [N]int{}
    cnt2 := [N]int{}

    mn1, mn2 := N, N 
    for _, x := range a {
        cnt1[x]++
        mn1 = min(mn1, x)
    }
    for _, x := range b {
        cnt2[x]++
        mn2 = min(mn2, x)
    }

    for i := range N {
        if cnt1[i] > 0 && cnt2[i] > 0 {
            return i
        }
    }

    return min(mn1 * 10 + mn2, mn2 * 10 + mn1)
}