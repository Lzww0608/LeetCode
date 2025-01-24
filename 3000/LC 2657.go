func findThePrefixCommonArray(a []int, b []int) []int {
    n := len(a)
    ans := make([]int, n)
    cnt := make([]int, n + 1)
    c := 0
    for i := 0; i < n; i++ {
        if cnt[a[i]]++; cnt[a[i]] == 2 {
            c++
        }
        if cnt[b[i]]++; cnt[b[i]] == 2 {
            c++
        }
        ans[i] = c
    }

    return ans
}