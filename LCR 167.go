func dismantlingAction(arr string) (ans int) {
    n := len(arr)
    l, r := 0, 0
    m := [128]bool{}
    for r < n {
        for l < r && m[arr[r]] {
            m[arr[l]] = false
            l++
        }
        m[arr[r]] = true
        ans = max(ans, r - l + 1)
        r++
    }

    return
}