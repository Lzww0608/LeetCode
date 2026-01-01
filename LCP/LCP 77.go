func runeReserve(a []int) int {
    sort.Ints(a)
    ans, cur := 1, 1
    for i := 1; i < len(a); i++ {
        if a[i] <= a[i - 1] + 1 {
            cur++
            ans = max(ans, cur)
        } else {
            cur = 1
        }
    }

    return ans
}