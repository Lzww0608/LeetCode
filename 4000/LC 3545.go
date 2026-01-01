func minDeletion(s string, k int) int {
    cnt := make([]int, 26)
    for _, c := range s {
        x := int(c - 'a')
        cnt[x]++
    }
    sort.Ints(cnt)
    ans := 0
    for i := range 26 - k {
        ans += cnt[i]
    }

    return ans
}